
# Confluence
./atlas --type confluence --action getPage --id "854950177"
./atlas --type confluence --action getPage --space "TEST" --title "Page A"
./atlas --type confluence --action getSpace --space "TEST"
./atlas --type confluence --action createPage --title "p1" --space "TEST" --body "lorem" --parent "136445988"
# create page with labels 'aa', 'bb'
./atlas --type confluence --action createPage --title "p1" --space "TEST" --body "lorem" --parent "136445988" --labels "aa,bb"
./atlas --type confluence --action copyPage --id "" --newTitle "" --parent "136125988"
./atlas --type confluence --action addLabel --id "858292123" --labels "label1,label2"

# Jira
./atlas --type jira --action getIssue --key "AAA-3"
./atlas --type jira --action getProject --name "AAA"
./atlas --type jira --action createIssue --summary "some" --descr "" --project "AAA"