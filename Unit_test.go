// package Lab5

// import (
// 	"testing"

// 	"github.com/asaskevich/govalidator"
// 	. "github.com/onsi/gomega"
// 	"gorm.io/gorm"
// )

// type Video struct {
// 	gorm.Model
// 	Name string `valid:"required~Name cannot be blank"`
// 	Url  string `gorm:"uniqueIndex" valid:"url"`
// }

// func TestUserNameNotBlank(t *testing.T) {
// 	g := NewGomegaWithT(t)

// 	user := Video{
// 		Name: "", //ผิด
// 		Url:  "https://www.google.com",
// 	}

// 	// ตรวจสอบด้วย govalidator
// 	ok, err := govalidator.ValidateStruct(user)

// 	// ok ต้องไม่เป็นค่า true แปลว่าต้อวจับ err ได้
// 	g.Expect(ok).ToNot(BeTrue())

// 	// ok ต้องไม่เป็นค่า nil แปลว่าต้อวจับ err ได้
// 	g.Expect(err).ToNot(BeNil())

// 	// err.Error ต้อวมี error message แสดงออกมา
// 	g.Expect(err.Error()).To(Equal("Name cannot be blank"))

// }

// func TestURLNotBlank(t *testing.T) {
// 	g := NewGomegaWithT(t)

// 	user := Video{
// 		Name: "Wallaya", //ผิด
// 		Url:  "asdf",
// 	}

// 	// ตรวจสอบด้วย govalidator
// 	ok, err := govalidator.ValidateStruct(user)

// 	// ok ต้องไม่เป็นค่า true แปลว่าต้อวจับ err ได้
// 	g.Expect(ok).ToNot(BeTrue())

// 	// ok ต้องไม่เป็นค่า nil แปลว่าต้อวจับ err ได้
// 	g.Expect(err).ToNot(BeNil())

// 	// err.Error ต้อวมี error message แสดงออกมา
// 	g.Expect(err.Error()).To(Equal("Url: asdf does not validate as url"))
// }

// func TestValidNotBlank(t *testing.T) {
// 	g := NewGomegaWithT(t)

// 	user := Video{
// 		Name: "Wallaya",
// 		Url:  "https://www.google.com/",
// 	}

// 	// ตรวจสอบด้วย govalidator
// 	ok, err := govalidator.ValidateStruct(user)

// 	// ok ต้องไม่เป็นค่า true แปลว่าต้อวจับ err ได้
// 	g.Expect(ok).To(BeTrue())

// 	// ok ต้องไม่เป็นค่า nil แปลว่าต้อวจับ err ได้
// 	g.Expect(err).To(BeNil())

// }


package Lab5


import (
	"testing"
	
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Name 	string		`valid:"required~Name cannot be blank"`
	Url		string		`gorm:"uniqueIndex" valid:"url"`
}

func TestVideoValidate(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Check data is valid", func(t *testing.T) {
		v := Video{
			Name: "Jakkrit Chaiwan",
			Url: "https://www.youtube.com/",
		}

		ok, err := govalidator.ValidateStruct(v)

		g.Expect(ok).To(BeTrue())

		g.Expect(err).To(BeNil())
	})
	t.Run("Check name is blank", func(t *testing.T) {
		v := Video{
			Name: "",
			Url: "https://www.youtube.com/",
		}

		ok, err := govalidator.ValidateStruct(v)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).ToNot(BeNil())

		g.Expect(err.Error()).To(Equal("Name cannot be blank"))

	})
	
	t.Run("Check url is valid", func(t *testing.T) {
		v := Video{
			Name: "Jakkrit",
			Url: "sdbvkwvslk",
		}

		ok, err := govalidator.ValidateStruct(v)

		g.Expect(ok).NotTo(BeTrue())

		g.Expect(err).ToNot(BeNil())

		g.Expect(err.Error()).To(Equal("Url: sdbvkwvslk does not validate as url"))

	})

	
}