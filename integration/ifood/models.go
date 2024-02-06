package ifood

type AuthRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}

type OrderDetails struct {
	StoreId            int         `json:"idLoja"`
	CustomerId         int         `json:"idCliente"`
	Customer           Customer    `json:"cliente" `
	OrderId            string      `json:"idPedido"`
	OrderCode          string      `json:"codigo"`
	CreatedAt          string      `json:"dataHora"`
	DeliveryAddress    Address     `json:"enderecoEntrega"`
	DeliveryFee        float32     `json:"valorEntrega"`
	TotalPrice         float32     `json:"valorTotal"`
	DiscountValue      float32     `json:"valorDesconto"`
	DocumentForTax     bool        `json:"cpfNaNota"`
	LineItems          []LineItems `json:"items"`
	ScheduledBeginDate *string     `json:"agendamentoDataInicio"`
	ScheduledEndDate   *string     `json:"agendamentoDataFim"`
	ScheduledBeginTime *string     `json:"agendamentoHoraInicio"`
	ScheduledEndTime   *string     `json:"agendamentoHoraFim"`
	Benefits           []Benefits  `json:"beneficios"`
}

type Customer struct {
	Id           int     `json:"id"`
	Name         string  `json:"nome"`
	Email        string  `json:"enderecoEntrega"`
	CpfDocument  *string `json:"cpf"`
	CnpjDocument *string `json:"cnpj"`
	Type         string  `json:"tipo"`
	Phone        string  `json:"telefoneCelular"`
}

type Address struct {
	Street        string  `json:"logradouro"`
	StreetNumber  string  `json:"numero"`
	Apartment     string  `json:"complemento"`
	Neighbourhood string  `json:"bairro"`
	City          string  `json:"cidade"`
	Province      string  `json:"estado"`
	ProvinceCode  string  `json:"uf"`
	Zip           string  `json:"cep"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
}

type LineItems struct {
	Id           *int     `json:"id"`
	Name         *string  `json:"produto"`
	BarCode      *string  `json:"codigoBarra"`
	InternalCode *string  `json:"plu"`
	UnitPrice    *float32 `json:"valor"`
	TotalPrice   *float32 `json:"valorTotal"`
	Quantity     *float32 `json:"quantidade"`
}

type Benefits struct {
	Sponsor *string  `json:"patrocinio"`
	Value   *float32 `json:"valor"`
	Type    *string  `json:"tipo"`
	ItemId  *int     `json:"itemId"`
}
