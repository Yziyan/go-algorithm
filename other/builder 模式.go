// @Author: Ciusyan 10/22/23

package other

import (
	"errors"
)

// builder 模式一般适用于构建较为复杂的对象，实现链式调用的效果

// ProductBuilder  用于构建 product
type ProductBuilder struct {
	p *product
}

// 构建一个 ProductBuilder 对象

func NewBuilder() *ProductBuilder {
	return &ProductBuilder{
		p: &product{},
	}
}

func (b *ProductBuilder) SetName(name string) *ProductBuilder {
	b.p.name = name
	return b
}

func (b *ProductBuilder) SetPrice(price float64) *ProductBuilder {
	b.p.price = price
	return b
}

func (b *ProductBuilder) SetQuantity(quantity int) *ProductBuilder {
	b.p.quantity = quantity
	return b
}

func (b *ProductBuilder) Build() *product {
	return b.p
}

// 这个对象甚至可以不对外暴露
type product struct {
	name     string
	price    float64
	quantity int
}

func (p *product) GetPrice() float64 {
	return p.price
}

// BuildV1 如果在链式调用的中间过程中，有可能有错误产生，那么我们还怎么使用链式调用呢？
// 如果是这副样子，外界只能先处理 err，再进行后续操作
//
//	p1, err := builder.SetName("zhiyan").SetPrice(3.20).SetQuantity(3).BuildV1()
//	require.NoError(t, err)
func (b *ProductBuilder) BuildV1() (*product, error) {
	if b.p == nil {
		return nil, errors.New("构建 Product 时有误")
	}

	return b.p, nil
}

func (p *product) GetPriceV1() (float64, error) {
	if p.price == 0.00 {
		return 0, errors.New("价格不能为 0.00")
	}

	return p.price, nil
}

// 但是其实我们如果还想要链式调用，我们可以设置一个中间状态的结果
type val struct {
	val *product
	err error
}

// BuildV2 返回中间结果，那么外界就可以继续使用链式调用了
func (b *ProductBuilder) BuildV2() *val {
	if b.p == nil {
		return &val{
			val: nil,
			err: errors.New("构建 Product 时有误"),
		}
	}

	return &val{
		val: b.p,
		err: nil,
	}
}

// GetPrice 这个调用，有可能出现错误，但是我们使用的是中间结果，
func (v *val) GetPrice() (float64, error) {
	if v.err != nil {
		return 0, v.err
	}

	if v.val.price == 0.00 {
		return 0, errors.New("价格不能为 0.00")
	}

	return v.val.price, nil
}

// MustProduct 如果我们想要一定有一个值，我们再中间也方便拦截
func (v *val) MustProduct() *product {
	if v.err != nil {
		panic(v.err)
	}

	return v.val
}
