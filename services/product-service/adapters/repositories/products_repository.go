package repositories

import (
	"github.com/zhunismp/nanow4t3r/services/product/core/domain"
	"gorm.io/gorm"
)

type ProductsRepositoryImpl struct {
	gormDB *gorm.DB
}

func NewProductsRepositoryImpl(gormDB *gorm.DB) *ProductsRepositoryImpl {
	return &ProductsRepositoryImpl{
		gormDB: gormDB,
	}
}

func (r *ProductsRepositoryImpl) GetAllProducts(activeOnly bool) ([]domain.BottledWater, error) {
	var products []domain.BottledWater
	query := r.gormDB

	if activeOnly {
		query = query.Where("is_active = ?", activeOnly)
	}

	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductsRepositoryImpl) GetProductByID(id int32) (domain.BottledWater, error) {
	var product domain.BottledWater
	if err := r.gormDB.First(&product, id).Error; err != nil {
		return domain.BottledWater{}, err
	}
	return product, nil
}

func (r *ProductsRepositoryImpl) CreateProduct(product domain.BottledWater) error {
	return r.gormDB.Create(&product).Error
}

func (r *ProductsRepositoryImpl) UpdateProduct(product domain.BottledWater) error {
	return r.gormDB.Save(&product).Error
}

func (r *ProductsRepositoryImpl) DeleteProductByID(id int32) error {
	return r.gormDB.Delete(&domain.BottledWater{}, id).Error
}
