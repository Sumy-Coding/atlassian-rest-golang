# Confluence REST service on Golang

The CLI app (or internal Go API) to interact with Confluence and Jira.

- Groovy version - https://github.com/AndriiMaliuta/confluence-rest-service-groovy
- Rust version - https://github.com/AndriiMaliuta/rust-atlas-rest

# Native CLI app
### Installation
Download execution file from Releases https://github.com/AndriiMaliuta/atlassian-rest-golang/releases.

add env args:

```bash
ATLAS_URL=
ATLAS_USER=
ATLAS_PASS=
```

---
### Confluence
```bash
atlas --type confluence --action getPage --id "854950177"
atlas --type confluence --action getPage --space "TEST" --title "Page A"
atlas --type confluence --action getSpace --space "TEST"
atlas --type confluence --action createPage --title "pp1" --space "test4" --body "lorem" --parent "858488858"
atlas --type confluence --action createSpace --name "s1" --category "test"
atlas --type confluence --action addLabels --id "111" --labels "aa,bb"

```
---
### Jira
```bash
atlas --type jira --action getIssue --key "AAA-3"
atlas --type jira --action createIssue --summary "" --description "" --project ""
atlas --type jira --action getProject --key "AD"
```
---
## Internal API examples
get services
```
grpcurl -plaintext localhost:9093 list
```

### Pages
getPage
```curl
./grpcurl -d '{"id": "123123"}' -plaintext localhost:50051 andmal.PageService.GetPage
```

Create 20 pages
```go
	for i := 2; i <= 20; i++ {
		space := spaceServ.GetSpace(cloudDcUrl, lToken, "TEST")
		pageServ.CreateContent(cloudDcUrl, lToken, "page", space.Key, space.Homepage.Id,
			fmt.Sprintf("GO page %d", i),
			ranServ.RandomString(100))
	}
```

```go
// 									=== ASYNC == several
	var waitG sync.WaitGroup
	for i := 11; i <= 30; i++ {
		waitG.Add(1)
		bod := ranServ.RandomString(100)
		go func(count int) {
			space := spaceServ.GetSpace(cloudDcUrl, dcToken, "DEV15")
			pageServ.CreateContent(cloudDcUrl, dcToken, "page", space.Key, space.Homepage.Id,
				fmt.Sprintf("GO page %d", count), bod)
		}(i)
	}
	waitG.Wait()
```

```go
// ========== Get Page
	fmt.Println(pageServ.GetPage(cloudDcUrl, lToken, "98383"))

	========== Create Page
	space := spaceServ.GetSpace(locUrl, lToken, "BBB")
	pageServ.CreateContent(locUrl, lToken, "page", space.Key, space.Homepage.Id, "GO page 1",
		ranServ.RandomString(100))
```

Create space(s)
```go 
fmt.Println(spaceServ.CreateSpace(locUrl, lToken, "DEV2", "DEV2"))

	for i := 3; i < 20; i++ {
		fmt.Println(spaceServ.CreateSpace(locUrl, lToken, fmt.Sprintf("DEV%d", i), fmt.Sprintf("DEV%d", i)))
	}

	// Operations took '124.458869' secs
	for a := 3; a <= 20; a++ {
		//wg.Add(1)
		go func() {
			key := fmt.Sprintf("DEV%d", a)
			sp := spaceServ.GetSpace(locUrl, lToken, key)
			for i := 40; i < 45; i++ {
				bod := ranServ.RandomString(15)
				pageServ.CreateContentAsync(locUrl, lToken, "page",
					sp.Key,
					sp.Homepage.Id,
					fmt.Sprintf("RST - %d", i), bod)
			}
		}()

	fmt.Println(spaceServ.CreateSpace(locUrl, lToken, fmt.Sprintf("DEV%d", a), fmt.Sprintf("DEV%d", a)))
	}
```

COMPLEX HIERARCHY
```go
	var count int
	for i := 40; i <= 50; i++ {
		space := spaceServ.GetSpace(locUrl, lToken, fmt.Sprintf("test%d", i))
		homePage := space.Homepage
		for i := 1; i <= 100; i++ {
			bod := ranServ.RandomString(100)
			pageServ.CreateContent(locUrl, lToken, "page", space.Key, homePage.Id, fmt.Sprintf("%s - %d", homePage.Title, i), bod)
		}
		count += i
	}
	
	log.Printf("%d pages created", count)

	// === GET space
	//fmt.Println(spaceServ.GetSpace(locUrl, lToken, "DEV"))

	// == Edit Page
	fmt.Println(pageServ.UpdatePage(locUrl, lToken, "2719745", "lorem", "lorem lorem lorem"))

	// == COPY page
	pageServ.CopyPage(locUrl, lToken, "65603", "2326554", true, true, false)

	// == COPY Hierarchy
	log.Println(pageServ.CopyPageDescs(bhtUrl, bhToken, "", "", "", true, true, false))

	// ==== ADD attach
	for _, att := range pageServ.GetPageAttaches(locUrl, lToken, "2719753").Results {
		fmt.Println(att.Id)
		pageServ.CopyAttach(locUrl, lToken, "2981909", att.Id)
	}

	// ==== Get attaches
	fmt.Println(pageServ.GetPageAttaches(locUrl, 1474565))

	// === GET attach
	fmt.Println(pageServ.GetAttach(locUrl, 1671169))

	// === Download attach
	pageServ.DownloadAttach(locUrl, 1671171)

	

	// === Comments
	pageServ.AddComment(locUrl, lToken, "2555915", "2555911")
```

### Labels
```go
fmt.Println(labServ.GetPageLabels(locUrl, lToken, "2719745"))

// ===== COPY labels
labels := []string{"aaa", "bbb"}
labServ.AddLabels(locUrl, lToken, "2555907", labels)
```
