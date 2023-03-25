package jira

//Get all workflows
//GET /rest/api/2/workflow

//Update property
//PUT /rest/api/2/workflow/{id}/

//Create property
//POST /rest/api/2/workflow/{id}/properties

//Get properties
//GET /rest/api/2/workflow/{id}/properties

//Delete property
//DELETE /rest/api/2/workflow/{id}/properties

type WorkflowService struct{}

func (ws WorkflowService) GetWorkflows(url string, token string) []Workflow {

	return nil
}

// workflowscheme

type WorkflowSchemeService struct{}

func (wss WorkflowSchemeService) GetWorkflowSchemes(url string, token string) []WorkflowScheme {

	return nil
}

//Create scheme
//POST /rest/api/2/workflowscheme

//Get by id
//GET /rest/api/2/workflowscheme/{id}

//Delete scheme
//DELETE /rest/api/2/workflowscheme/{id}

//Update
//PUT /rest/api/2/workflowscheme/{id}
