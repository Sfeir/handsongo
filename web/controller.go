package web

import (
	"github.com/Sfeir/handsongo/dao"
	"github.com/Sfeir/handsongo/model"
	logger "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"net/http"
	"strconv"
)

const (
	prefix = "/v1/spirits"
)

// SpiritController is a handler of spirits
type SpiritController struct {
	spiritDao dao.SpiritDAO
	Routes    []Route
	Prefix    string
}

// NewSpiritController creates a new spirit handler to manage spirits
func NewSpiritController(spiritDAO dao.SpiritDAO) *SpiritController {
	controller := SpiritController{
		spiritDao: spiritDAO,
		Prefix:    prefix,
	}

	// build routes
	routes := []Route{}
	// GetAll
	routes = append(routes, Route{
		Name:        "Get all spirits",
		Method:      http.MethodGet,
		Pattern:     "",
		HandlerFunc: controller.GetAll,
	})
	// Get
	routes = append(routes, Route{
		Name:        "Get one spirit",
		Method:      http.MethodGet,
		Pattern:     "/{id}",
		HandlerFunc: controller.Get,
	})
	// Create
	routes = append(routes, Route{
		Name:        "Create a spirit",
		Method:      http.MethodPost,
		Pattern:     "",
		HandlerFunc: controller.Create,
	})
	// Update
	routes = append(routes, Route{
		Name:        "Update a spirit",
		Method:      http.MethodPut,
		Pattern:     "/{id}",
		HandlerFunc: controller.Update,
	})
	// Delete
	routes = append(routes, Route{
		Name:        "Delete a spirit",
		Method:      http.MethodDelete,
		Pattern:     "/{id}",
		HandlerFunc: controller.Delete,
	})

	controller.Routes = routes

	return &controller
}

// GetAll retrieve all entities with optional paging of items (start / end are item counts 50 to 100 for example)
func (sh *SpiritController) GetAll(w http.ResponseWriter, r *http.Request) {

	startStr := ParamAsString("start", r)
	endStr := ParamAsString("end", r)

	start := dao.NoPaging
	end := dao.NoPaging
	var err error
	if startStr != "" && endStr != "" {
		start, err = strconv.Atoi(startStr)
		if err != nil {
			start = dao.NoPaging
		}
		end, err = strconv.Atoi(endStr)
		if err != nil {
			end = dao.NoPaging
		}
	}

	// find all spirits
	spirits, err := sh.spiritDao.GetAllSpirits(start, end)
	if err != nil {
		logger.WithField("error", err).Warn("unable to retrieve spirits")
		SendJSONError(w, "Error while retrieving spirits", http.StatusInternalServerError)
		return
	}

	logger.WithField("spirits", spirits).Debug("spirits found")
	SendJSONOk(w, spirits)
}

// Get retrieve an entity by id
func (sh *SpiritController) Get(w http.ResponseWriter, r *http.Request) {
	// get the spirit ID from the URL
	spiritID := ParamAsString("id", r)

	// find spirit
	spirit, err := sh.spiritDao.GetSpiritByID(spiritID)
	if err != nil {
		if err == mgo.ErrNotFound {
			logger.WithField("error", err).WithField("spirit ID", spiritID).Warn("unable to retrieve spirit by ID")
			SendJSONNotFound(w)
			return
		}

		logger.WithField("error", err).WithField("spirit ID", spiritID).Warn("unable to retrieve spirit by ID")
		SendJSONError(w, "Error while retrieving spirit by ID", http.StatusInternalServerError)
		return
	}

	logger.WithField("spirits", spirit).Debug("spirit found")
	SendJSONOk(w, spirit)
}

// Create create an entity
func (sh *SpiritController) Create(w http.ResponseWriter, r *http.Request) {
	// spirit to be created
	spirit := &model.Spirit{}
	// get the content body
	err := GetJSONContent(spirit, r)

	if err != nil {
		logger.WithField("error", err).Warn("unable to decode spirit to create")
		SendJSONError(w, "Error while decoding spirit to create", http.StatusBadRequest)
		return
	}

	// save spirit
	err = sh.spiritDao.SaveSpirit(spirit)
	if err != nil {
		logger.WithField("error", err).WithField("spirit", *spirit).Warn("unable to create spirit")
		SendJSONError(w, "Error while creating spirit", http.StatusInternalServerError)
		return
	}

	// send response
	SendJSONOk(w, spirit)
}

// Update update an entity by id
func (sh *SpiritController) Update(w http.ResponseWriter, r *http.Request) {
	// get the spirit ID from the URL
	spiritID := ParamAsString("id", r)

	// spirit to be created
	spirit := &model.Spirit{}
	// get the content body
	err := GetJSONContent(spirit, r)

	if err != nil {
		logger.WithField("error", err).Warn("unable to decode spirit to create")
		SendJSONError(w, "Error while decoding spirit to create", http.StatusBadRequest)
		return
	}

	// save spirit
	_, err = sh.spiritDao.UpsertSpirit(spiritID, spirit)
	if err != nil {
		logger.WithField("error", err).WithField("spirit", *spirit).Warn("unable to create spirit")
		SendJSONError(w, "Error while creating spirit", http.StatusInternalServerError)
		return
	}

	// send response
	SendJSONOk(w, spirit)
}

// Delete delete an entity by id
func (sh *SpiritController) Delete(w http.ResponseWriter, r *http.Request) {
	// get the spirit ID from the URL
	spiritID := ParamAsString("id", r)

	// find spirit
	err := sh.spiritDao.DeleteSpirit(spiritID)
	if err != nil {
		logger.WithField("error", err).WithField("spirit ID", spiritID).Warn("unable to delete spirit by ID")
		SendJSONError(w, "Error while deleting spirit by ID", http.StatusInternalServerError)
		return
	}

	logger.WithField("spiritID", spiritID).Debug("spirit deleted")
	SendJSONOk(w, nil)
}
