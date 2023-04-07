package zabbixgosdk

import "testing"

const (
	imageName          = "test-image"
	updatedImageName   = "test-image-update"
	base64Image        = "iVBORw0KGgoAAAANSUhEUgAAAIAAAAB5CAIAAAB+wb5NAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAAEnQAABJ0Ad5mH3gAAAUMSURBVHhe7Z1LchsxDERzLh9I59FpdBkfxiHAzwAgSI4TOZ1Fv4VLg08DZJdKWU1+fREoNAAMDQBDA8DQADA0AAwNAEMDwNAAMDQADA0AQwPA0AAwNAAMDQBDA8DQADA0AAwNALMw4PX4NfF4taTw+fyYYoXQ9/H8bIljqippkROViK2daJsM9ivFrMdqXZVh87CO9JiISmwX9tz6BsyiEnk8xr0N3HX52zyk2kMY5esmND0tFtpNfq43+NGltH+WrqtJy8Kz7drtm3DHAL+AIEPLGEn4acmuveCQGjqaqYWuaMIU5sx77yTXmUnHlcpDy0rhet+cswGJqszUiLs5we9qdjumjEyf6Dsi7hYy8vZ1Wx87kQiZUD+JCq/XXXEyINvKXJ6kbdbt6hc/pNwMCQhXycTx/lcFfrRFOwqxKekw4vU2nrLxbp0VewPSQ9SJdZ9Y0G+uY/Y+pPwU1Y2ndvjJTrxG090LyXUa6uBLRUg6jHjW8Q12Bqj0tKtfJ9Qkyb7XIWW3r6If5c8VrQ2NEtZn21S5pBYFY48oaGm5FvWbK1LQQu2jtoSqO2wMSOYWJDoxymKPrlXPcUiZK+iVscMj2XhxTipvN2O3bDYvmDFSqFntWO+7YGlAfr7VMj0U05tjhNQYZer2R0qzVipOFMzUPZvNvYg82H1vqV8sDFjtmZ3JBn1eVfrjIdWGhcGrPSqaNaIFI9XzQW0lVzqNkuj0ymTzIHplbd8tcgPqkMjj5VcZXGGdb4hHMoSULq1z/QBt250o6Mb1fHp7Ne7Um829iDfgxsKe3Y8w+QfQADA0AAwNAEMDwNAAMDQADA0AQwPA0AAwNAAMDQBDA8DQADA0AAwNAEMDwNAAMDQADA0AQwPA0AAwNAAMDQBDA8DQADA0AAwNAEMDwNAAMDQADA0AQwPA0AAwNAAMDQBDA8DQADA0AAwNAEMDwNAAMDQADA0AQwPA0AAwNAAMDQBDA8DQADA0AAwNAJMbEN5S15FX0cU32F0vqLPvvKuxubbSdUJnf1Q05iLhrXqFkG4kynllS8bX/GXbXKx3WN/MjsM3wB6mIpG4ckH3GvFS5GtynR4ZhzIlGjPPUu9FpSLZZFLWV/AmhVInuZCSluzN/IX9DqqX7LPnTQbE64rkOu6aHg93T05RH+4eblJ+vvSv715lJF6654WPO+Q3c+Jd3wAJh0pLrtMj9TLq65dtrD+4xJGTsmIS7jQS0MK48SwRkY7kZg68y4C64NKEXKdHxjUYcXPg89kdmXKU6DVjcsM8+5Vv7GCW/wZ/ZIDFZnXLKSrkOj1yHfs6pzmxP7xbwGtWcmWJjusZJVdesY9+6o0dXLAwqre87RtwobuGtlynR+y59RglYU7sDz+YNSsLZQnXjyZoJ9sSQcf2xxs7+O67/IABhWndXKdH3DW0c8vPYy+Q2tBemDUrK+W+lORtbEzWMZGePe9gde/z/xnQut3/31AjYeqsWVkq163cvzBtPjmZDR13SPpv8CYDStQEpca35To94q5JUAGnocefi7xmZa1cRaJsy0ub0xdc8LBDKnDkr3+E+9B+OCU9yKyTX5PQ1HxLnL047U5Zcvb5ysdMYwqvdwiZQqI3cTCA/DQ0AAwNAEMDwNAAMDQADA0AQwPA0AAwNAAMDQBDA8DQADA0AAwNAEMDwNAAMDQADA0AQwPA0AAwNAAMDQBDA8DQADA0AAwNAEMDoHx9/QYhko0uvFWPCwAAAABJRU5ErkJggg=="
	updatedBase64Image = "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"
)

var imageId string

func TestImageCreate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	i, err := client.Image.Create(&ImageCreateParameters{
		Name:      imageName,
		ImageType: ImageIcon,
		Base64:    base64Image,
	})

	if err != nil {
		t.Fatalf("Error when creating image '%s'.\nReason : %v", iconMapName, err)
	}

	if i == nil || len(i.ImageIds) == 0 {
		t.Fatalf("Create method should returned a list with the id of the image.\nAn empty list was returned.")
	}

	imageId = i.ImageIds[0]
}

func TestImageGet(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	i, err := client.Image.Get(&ImageGetParameters{
		ImageIds: []string{
			imageId,
		},
	})

	if err != nil {
		t.Fatalf("Error when getting image '%s'.\nReason : %v", imageId, err)
	}

	if len(i) == 0 {
		t.Fatalf("Get method should returned a list of images matching the given criteria.\nAn empty list was returned.")
	}

	if i[0].ImageId != imageId {
		t.Fatalf("Wrong image returned.\nExpected : %s\nReturned : %s", imageId, i[0].ImageId)
	}
}

func TestImageUpdate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	i, err := client.Image.Update(&ImageUpdateParameters{
		ImageId: imageId,
		Name:    updatedImageName,
		Base64:  updatedBase64Image,
	})

	if err != nil {
		t.Fatalf("Error when updating image '%s'.\nReason : %v", imageId, err)
	}

	if i == nil || len(i.ImageIds) == 0 {
		t.Fatalf("Update method should returned a list of update images.\nAn empty list was returned.")
	}

	if i.ImageIds[0] != imageId {
		t.Fatalf("Wrong image returned.\nExpected : %s\nReturned : %s", imageId, i.ImageIds[0])
	}
}

func TestImageDelete(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	i, err := client.Image.Delete([]string{
		imageId,
	})

	if err != nil {
		t.Fatalf("Error when deleting image '%s'.\nReason : %v", imageId, err)
	}

	if i.ImageIds == nil || len(i.ImageIds) == 0 {
		t.Fatalf("Delete method should returned a list with the id of the deleted images.\nAn empty list was returned.")
	}

	if i.ImageIds[0] != imageId {
		t.Fatalf("Wrong id returned.\nExpected : %s\nReturned : %s", imageId, i.ImageIds[0])
	}
}
