package keto

import (
	"net/http"

	"github.com/sdslabs/nymeria/config"
)

var (
	CreateRelationshipEndpoint = Endpoint{
		URL:    config.KetoWriteURL + "/admin/relation-tuples",
		Method: http.MethodPut,
	}
	QueryRelationshipsEndpoint = Endpoint{
		URL:    config.KetoReadURL + "/relation-tuples",
		Method: http.MethodGet,
	}
	DeleteRelationshipsEndpoint = Endpoint{
		URL:    config.KetoWriteURL + "/admin/relation-tuples",
		Method: http.MethodDelete,
	}

	CheckPermissionEndpoint = Endpoint{
		URL:    config.KetoReadURL + "/relation-tuples/check",
		Method: http.MethodPost,
	}
)
