package orders

type Order struct {
	StoreId            int         `bson:"idLoja"`
	CustomerId         int         `bson:"idCliente"`
	Customer           Customer    `bson:"cliente" `
	OrderId            string      `bson:"idPedido"`
	OrderCode          string      `bson:"codigo"`
	CreatedAt          string      `bson:"dataHora"`
	DeliveryAddress    Address     `bson:"enderecoEntrega"`
	DeliveryFee        float32     `bson:"valorEntrega"`
	TotalPrice         float32     `bson:"valorTotal"`
	DiscountValue      float32     `bson:"valorDesconto"`
	DocumentForTax     bool        `bson:"cpfNaNota"`
	LineItems          []LineItems `bson:"items"`
	ScheduledBeginDate *string     `bson:"agendamentoDataInicio"`
	ScheduledEndDate   *string     `bson:"agendamentoDataFim"`
	ScheduledBeginTime *string     `bson:"agendamentoHoraInicio"`
	ScheduledEndTime   *string     `bson:"agendamentoHoraFim"`
	Benefits           []Benefits  `bson:"beneficios"`
}

type Customer struct {
	Id           int     `bson:"id"`
	Name         string  `bson:"nome"`
	Email        string  `bson:"enderecoEntrega"`
	CpfDocument  *string `bson:"cpf"`
	CnpjDocument *string `bson:"cnpj"`
	Type         string  `bson:"tipo"`
	Phone        string  `bson:"telefoneCelular"`
}

type Address struct {
	Street        string  `bson:"logradouro"`
	StreetNumber  string  `bson:"numero"`
	Apartment     string  `bson:"complemento"`
	Neighbourhood string  `bson:"bairro"`
	City          string  `bson:"cidade"`
	Province      string  `bson:"estado"`
	ProvinceCode  string  `bson:"uf"`
	Zip           string  `bson:"cep"`
	Latitude      float64 `bson:"latitude"`
	Longitude     float64 `bson:"longitude"`
}

type LineItems struct {
	Id           *int     `bson:"id"`
	Name         *string  `bson:"produto"`
	BarCode      *string  `bson:"codigoBarra"`
	InternalCode *string  `bson:"plu"`
	UnitPrice    *float32 `bson:"valor"`
	TotalPrice   *float32 `bson:"valorTotal"`
	Quantity     *float32 `bson:"quantidade"`
}

type Benefits struct {
	Sponsor *string  `bson:"patrocinio"`
	Value   *float32 `bson:"valor"`
	Type    *string  `bson:"tipo"`
	ItemId  *int     `bson:"itemId"`
}
