package version

import (
	"base/application/dto"
	"base/common/response"
	"base/infrastructure/provider"
	"base/infrastructure/svc"
	"encoding/json"
	"net/http"
)

func DeleteVersionHandler(svc *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.VersionDeleteRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			response.Response(r, w, nil, err)
			return
		}
		versionService := provider.InitializeVersionService(svc)
		resp, err := versionService.Delete(r.Context(), &req)
		response.Response(r, w, resp, err)
	}
}
