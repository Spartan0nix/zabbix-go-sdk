package zabbixgosdk

// ImageService create a new service to access image related methods and functions.
type ImageService struct {
	Client *ApiClient
}

// ImageType define the available types of image
type ImageType string

const (
	ImageIcon       ImageType = "1"
	ImageBackGround ImageType = "2"
)

// Image properties.
type Image struct {
	// ReadOnly
	ImageId   string    `json:"imageid"`
	Name      string    `json:"name"`
	ImageType ImageType `json:"imagetype"`
	Base64    string    `json:"image"`
}

// ImageResponse define the server response format for Image methods.
type ImageResponse struct {
	ImageIds []string `json:"imageids"`
}

// ImageCreateParameters define the properties needed to create a new Image
type ImageCreateParameters struct {
	Name      string    `json:"name"`
	ImageType ImageType `json:"imagetype"`
	// Base64 enconded image
	// Default maximum size : 1 MB
	// Support format : PNG, JPEG, GIF
	Base64 string `json:"image"`
}

// Create is used to create a new Image.
func (i *ImageService) Create(p *ImageCreateParameters) (*ImageResponse, error) {
	req := i.Client.NewRequest("image.create", p)

	res, err := i.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := ImageResponse{}
	err = i.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// Delete is used to delete one or multiples Images.
func (i *ImageService) Delete(ids []string) (*ImageResponse, error) {
	req := i.Client.NewRequest("image.delete", ids)

	res, err := i.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := ImageResponse{}
	err = i.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// ImageGetParameters define the properties used to search Images.
// Properties using the 'omitempty' json parameters are optional
type ImageGetParameters struct {
	ImageIds               []string    `json:"imageids,omitempty"`
	Sysmapids              []string    `json:"sysmapids,omitempty"`
	SelectImage            interface{} `json:"select_image,omitempty"`
	CountOutput            bool        `json:"countOutput,omitempty"`
	Editable               bool        `json:"editable,omitempty"`
	ExcludeSearch          bool        `json:"excludeSearch,omitempty"`
	Filter                 interface{} `json:"filter,omitempty"`
	Limit                  string      `json:"limit,omitempty"`
	Output                 interface{} `json:"output,omitempty"`
	PreserveKeys           bool        `json:"preservekeys,omitempty"`
	Search                 interface{} `json:"search,omitempty"`
	SearchByAny            bool        `json:"searchByAny,omitempty"`
	SearchWildcardsEnabled bool        `json:"searchWildcardsEnabled,omitempty"`
	SortField              []string    `json:"sortfield,omitempty"`
	SortOrder              []string    `json:"sortorder,omitempty"`
	StartSearch            bool        `json:"startSearch,omitempty"`
}

// Get is used to retrieve one or multiples Images matching the given criteria(s).
func (i *ImageService) Get(p *ImageGetParameters) ([]*Image, error) {
	req := i.Client.NewRequest("image.get", p)

	res, err := i.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := make([]*Image, 0)
	err = i.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// ImageUpdateParameters define the properties used to update Images.
// Properties using the 'omitempty' json parameters are optional
type ImageUpdateParameters struct {
	ImageId   string    `json:"imageid"`
	Name      string    `json:"name,omitempty"`
	ImageType ImageType `json:"imagetype,omitempty"`
	Base64    string    `json:"image,omitempty"`
}

// Update is used to retrieve one or multiples Images matching the given criteria(s).
func (i *ImageService) Update(p *ImageUpdateParameters) (*ImageResponse, error) {
	req := i.Client.NewRequest("image.update", p)

	res, err := i.Client.Post(req)
	if err != nil {
		return nil, err
	}

	r := ImageResponse{}
	err = i.Client.ConvertResponse(*res, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
