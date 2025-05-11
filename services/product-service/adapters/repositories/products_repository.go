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

func (r *ProductsRepositoryImpl) GetAllProducts(activeOnly bool) ([]domain.Product, error) {
	var products []domain.Product
	query := r.gormDB

	if activeOnly {
		query = query.Where("is_active = ?", true)
	}

	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductsRepositoryImpl) GetProductByID(id uint32) (domain.Product, error) {
	var product domain.Product
	if err := r.gormDB.First(&product, id).Error; err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (r *ProductsRepositoryImpl) CreateProduct(product domain.Product) error {
	return r.gormDB.Create(&product).Error
}

func (r *ProductsRepositoryImpl) UpdateProduct(product domain.Product) error {
	return r.gormDB.Save(&product).Error
}

func (r *ProductsRepositoryImpl) DeleteProductByID(id uint32) error {
	return r.gormDB.Delete(&domain.Product{}, id).Error
}
