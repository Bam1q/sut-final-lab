package entity

import (
	"testing"

	"github.com/Bam1q/sut-final-lab/backend/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNagative(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`Test Positive`, func(t *testing.T) {
		device := entity.Devices{
			DeviceID:   "1234567899", // ผิด
			IpAddress:  "192.168.160.1",
			DeviceName: "Gomaga",
		}

		ok, err := govalidator.ValidateStruct(device)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expext(err.Error()).To(Equal("DeviceID must start with 2 uppercase letters followed by 10 digits"))

	})
}
