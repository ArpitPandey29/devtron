package resourceQualifiers

import (
	"github.com/devtron-labs/devtron/pkg/devtronResource"
	"github.com/devtron-labs/devtron/pkg/sql"
	"github.com/go-pg/pg"
	"go.uber.org/zap"
)

type QualifierMappingService interface {
	CreateQualifierMappings(qualifierMappings []*QualifierMapping, tx *pg.Tx) ([]*QualifierMapping, error)
	GetQualifierMappings(resourceType ResourceType, scope *Scope, resourceIds []int) ([]*QualifierMapping, error)
	DeleteAllQualifierMappings(resourceType ResourceType, auditLog sql.AuditLog, tx *pg.Tx) error
}

type QualifierMappingServiceImpl struct {
	logger                     *zap.SugaredLogger
	qualifierMappingRepository QualifiersMappingRepository
	devtronResourceService     devtronResource.DevtronResourceService
}

func NewQualifierMappingServiceImpl(logger *zap.SugaredLogger, qualifierMappingRepository QualifiersMappingRepository, devtronResourceService devtronResource.DevtronResourceService) (*QualifierMappingServiceImpl, error) {
	return &QualifierMappingServiceImpl{
		logger:                     logger,
		qualifierMappingRepository: qualifierMappingRepository,
		devtronResourceService:     devtronResourceService,
	}, nil
}

func (impl QualifierMappingServiceImpl) CreateQualifierMappings(qualifierMappings []*QualifierMapping, tx *pg.Tx) ([]*QualifierMapping, error) {
	return impl.qualifierMappingRepository.CreateQualifierMappings(qualifierMappings, tx)
}

func (impl QualifierMappingServiceImpl) GetQualifierMappings(resourceType ResourceType, scope *Scope, resourceIds []int) ([]*QualifierMapping, error) {
	searchableKeyNameIdMap := impl.devtronResourceService.GetAllSearchableKeyNameIdMap()
	return impl.qualifierMappingRepository.GetQualifierMappings(resourceType, scope, searchableKeyNameIdMap, resourceIds)
}

func (impl QualifierMappingServiceImpl) DeleteAllQualifierMappings(resourceType ResourceType, auditLog sql.AuditLog, tx *pg.Tx) error {
	return impl.qualifierMappingRepository.DeleteAllQualifierMappings(resourceType, auditLog, tx)
}
