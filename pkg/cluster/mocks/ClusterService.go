// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import (
	context "context"

	cluster "github.com/devtron-labs/devtron/pkg/cluster"

	mock "github.com/stretchr/testify/mock"

	repository "github.com/devtron-labs/devtron/pkg/cluster/repository"

	userrepository "github.com/devtron-labs/devtron/pkg/user/repository"

	util "github.com/devtron-labs/devtron/internal/util"

	v1 "k8s.io/client-go/kubernetes/typed/core/v1"

	v1alpha1 "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
)

// ClusterService is an autogenerated mock type for the ClusterService type
type ClusterService struct {
	mock.Mock
}

// ConnectClustersInBatch provides a mock function with given fields: clusters, clusterExistInDb
func (_m *ClusterService) ConnectClustersInBatch(clusters []*cluster.ClusterBean, clusterExistInDb bool) {
	_m.Called(clusters, clusterExistInDb)
}

// ConvertClusterBeanObjectToCluster provides a mock function with given fields: bean
func (_m *ClusterService) ConvertClusterBeanObjectToCluster(bean *cluster.ClusterBean) *v1alpha1.Cluster {
	ret := _m.Called(bean)

	var r0 *v1alpha1.Cluster
	if rf, ok := ret.Get(0).(func(*cluster.ClusterBean) *v1alpha1.Cluster); ok {
		r0 = rf(bean)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Cluster)
		}
	}

	return r0
}

// ConvertClusterBeanToCluster provides a mock function with given fields: clusterBean, userId
func (_m *ClusterService) ConvertClusterBeanToCluster(clusterBean *cluster.ClusterBean, userId int32) *repository.Cluster {
	ret := _m.Called(clusterBean, userId)

	var r0 *repository.Cluster
	if rf, ok := ret.Get(0).(func(*cluster.ClusterBean, int32) *repository.Cluster); ok {
		r0 = rf(clusterBean, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repository.Cluster)
		}
	}

	return r0
}

// CreateGrafanaDataSource provides a mock function with given fields: clusterBean, env
func (_m *ClusterService) CreateGrafanaDataSource(clusterBean *cluster.ClusterBean, env *repository.Environment) (int, error) {
	ret := _m.Called(clusterBean, env)

	var r0 int
	if rf, ok := ret.Get(0).(func(*cluster.ClusterBean, *repository.Environment) int); ok {
		r0 = rf(clusterBean, env)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*cluster.ClusterBean, *repository.Environment) error); ok {
		r1 = rf(clusterBean, env)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: bean, userId
func (_m *ClusterService) Delete(bean *cluster.ClusterBean, userId int32) error {
	ret := _m.Called(bean, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(*cluster.ClusterBean, int32) error); ok {
		r0 = rf(bean, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteFromDb provides a mock function with given fields: bean, userId
func (_m *ClusterService) DeleteFromDb(bean *cluster.ClusterBean, userId int32) error {
	ret := _m.Called(bean, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(*cluster.ClusterBean, int32) error); ok {
		r0 = rf(bean, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchRolesFromGroup provides a mock function with given fields: userId
func (_m *ClusterService) FetchRolesFromGroup(userId int32) ([]*userrepository.RoleModel, error) {
	ret := _m.Called(userId)

	var r0 []*userrepository.RoleModel
	if rf, ok := ret.Get(0).(func(int32) []*userrepository.RoleModel); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*userrepository.RoleModel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int32) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields:
func (_m *ClusterService) FindAll() ([]*cluster.ClusterBean, error) {
	ret := _m.Called()

	var r0 []*cluster.ClusterBean
	if rf, ok := ret.Get(0).(func() []*cluster.ClusterBean); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*cluster.ClusterBean)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllActive provides a mock function with given fields:
func (_m *ClusterService) FindAllActive() ([]cluster.ClusterBean, error) {
	ret := _m.Called()

	var r0 []cluster.ClusterBean
	if rf, ok := ret.Get(0).(func() []cluster.ClusterBean); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cluster.ClusterBean)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllForAutoComplete provides a mock function with given fields:
func (_m *ClusterService) FindAllForAutoComplete() ([]cluster.ClusterBean, error) {
	ret := _m.Called()

	var r0 []cluster.ClusterBean
	if rf, ok := ret.Get(0).(func() []cluster.ClusterBean); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cluster.ClusterBean)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllForClusterByUserId provides a mock function with given fields: userId, isActionUserSuperAdmin
func (_m *ClusterService) FindAllForClusterByUserId(userId int32, isActionUserSuperAdmin bool) ([]cluster.ClusterBean, error) {
	ret := _m.Called(userId, isActionUserSuperAdmin)

	var r0 []cluster.ClusterBean
	if rf, ok := ret.Get(0).(func(int32, bool) []cluster.ClusterBean); ok {
		r0 = rf(userId, isActionUserSuperAdmin)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cluster.ClusterBean)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int32, bool) error); ok {
		r1 = rf(userId, isActionUserSuperAdmin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllNamespacesByUserIdAndClusterId provides a mock function with given fields: userId, clusterId, isActionUserSuperAdmin
func (_m *ClusterService) FindAllNamespacesByUserIdAndClusterId(userId int32, clusterId int, isActionUserSuperAdmin bool) ([]string, error) {
	ret := _m.Called(userId, clusterId, isActionUserSuperAdmin)

	var r0 []string
	if rf, ok := ret.Get(0).(func(int32, int, bool) []string); ok {
		r0 = rf(userId, clusterId, isActionUserSuperAdmin)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int32, int, bool) error); ok {
		r1 = rf(userId, clusterId, isActionUserSuperAdmin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllWithoutConfig provides a mock function with given fields:
func (_m *ClusterService) FindAllWithoutConfig() ([]*cluster.ClusterBean, error) {
	ret := _m.Called()

	var r0 []*cluster.ClusterBean
	if rf, ok := ret.Get(0).(func() []*cluster.ClusterBean); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*cluster.ClusterBean)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: id
func (_m *ClusterService) FindById(id int) (*cluster.ClusterBean, error) {
	ret := _m.Called(id)

	var r0 *cluster.ClusterBean
	if rf, ok := ret.Get(0).(func(int) *cluster.ClusterBean); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cluster.ClusterBean)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByIdWithoutConfig provides a mock function with given fields: id
func (_m *ClusterService) FindByIdWithoutConfig(id int) (*cluster.ClusterBean, error) {
	ret := _m.Called(id)

	var r0 *cluster.ClusterBean
	if rf, ok := ret.Get(0).(func(int) *cluster.ClusterBean); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cluster.ClusterBean)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByIds provides a mock function with given fields: id
func (_m *ClusterService) FindByIds(id []int) ([]cluster.ClusterBean, error) {
	ret := _m.Called(id)

	var r0 []cluster.ClusterBean
	if rf, ok := ret.Get(0).(func([]int) []cluster.ClusterBean); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cluster.ClusterBean)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOne provides a mock function with given fields: clusterName
func (_m *ClusterService) FindOne(clusterName string) (*cluster.ClusterBean, error) {
	ret := _m.Called(clusterName)

	var r0 *cluster.ClusterBean
	if rf, ok := ret.Get(0).(func(string) *cluster.ClusterBean); ok {
		r0 = rf(clusterName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cluster.ClusterBean)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(clusterName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOneActive provides a mock function with given fields: clusterName
func (_m *ClusterService) FindOneActive(clusterName string) (*cluster.ClusterBean, error) {
	ret := _m.Called(clusterName)

	var r0 *cluster.ClusterBean
	if rf, ok := ret.Get(0).(func(string) *cluster.ClusterBean); ok {
		r0 = rf(clusterName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cluster.ClusterBean)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(clusterName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllClusterNamespaces provides a mock function with given fields:
func (_m *ClusterService) GetAllClusterNamespaces() map[string][]string {
	ret := _m.Called()

	var r0 map[string][]string
	if rf, ok := ret.Get(0).(func() map[string][]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]string)
		}
	}

	return r0
}

// GetClusterConfig provides a mock function with given fields: _a0
func (_m *ClusterService) GetClusterConfig(_a0 *cluster.ClusterBean) (*util.ClusterConfig, error) {
	ret := _m.Called(_a0)

	var r0 *util.ClusterConfig
	if rf, ok := ret.Get(0).(func(*cluster.ClusterBean) *util.ClusterConfig); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*util.ClusterConfig)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*cluster.ClusterBean) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetK8sClient provides a mock function with given fields:
func (_m *ClusterService) GetK8sClient() (*v1.CoreV1Client, error) {
	ret := _m.Called()

	var r0 *v1.CoreV1Client
	if rf, ok := ret.Get(0).(func() *v1.CoreV1Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.CoreV1Client)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HandleErrorInClusterConnections provides a mock function with given fields: clusters, respMap, clusterExistInDb
func (_m *ClusterService) HandleErrorInClusterConnections(clusters []*cluster.ClusterBean, respMap map[int]error, clusterExistInDb bool) {
	_m.Called(clusters, respMap, clusterExistInDb)
}

// Save provides a mock function with given fields: parent, bean, userId
func (_m *ClusterService) Save(parent context.Context, bean *cluster.ClusterBean, userId int32) (*cluster.ClusterBean, error) {
	ret := _m.Called(parent, bean, userId)

	var r0 *cluster.ClusterBean
	if rf, ok := ret.Get(0).(func(context.Context, *cluster.ClusterBean, int32) *cluster.ClusterBean); ok {
		r0 = rf(parent, bean, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cluster.ClusterBean)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *cluster.ClusterBean, int32) error); ok {
		r1 = rf(parent, bean, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, bean, userId
func (_m *ClusterService) Update(ctx context.Context, bean *cluster.ClusterBean, userId int32) (*cluster.ClusterBean, error) {
	ret := _m.Called(ctx, bean, userId)

	var r0 *cluster.ClusterBean
	if rf, ok := ret.Get(0).(func(context.Context, *cluster.ClusterBean, int32) *cluster.ClusterBean); ok {
		r0 = rf(ctx, bean, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cluster.ClusterBean)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *cluster.ClusterBean, int32) error); ok {
		r1 = rf(ctx, bean, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateKubeconfig provides a mock function with given fields: kubeConfig
func (_m *ClusterService) ValidateKubeconfig(kubeConfig string) (map[string]*cluster.ValidateClusterBean, error) {
	ret := _m.Called(kubeConfig)

	var r0 map[string]*cluster.ValidateClusterBean
	if rf, ok := ret.Get(0).(func(string) map[string]*cluster.ValidateClusterBean); ok {
		r0 = rf(kubeConfig)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*cluster.ValidateClusterBean)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(kubeConfig)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewClusterService interface {
	mock.TestingT
	Cleanup(func())
}

// NewClusterService creates a new instance of ClusterService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClusterService(t mockConstructorTestingTNewClusterService) *ClusterService {
	mock := &ClusterService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
