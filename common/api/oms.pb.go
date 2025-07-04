// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: api/oms.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Order struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ID            string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CustomerID    string                 `protobuf:"bytes,2,opt,name=CustomerID,proto3" json:"CustomerID,omitempty"`
	Status        string                 `protobuf:"bytes,3,opt,name=Status,proto3" json:"Status,omitempty"`
	Items         []*Item                `protobuf:"bytes,4,rep,name=Items,proto3" json:"Items,omitempty"`
	PaymentLink   string                 `protobuf:"bytes,5,opt,name=PaymentLink,proto3" json:"PaymentLink,omitempty"`
	Image         string                 `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"`
	ResultsBase64 []string               `protobuf:"bytes,7,rep,name=results_base64,json=resultsBase64,proto3" json:"results_base64,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_api_oms_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_api_oms_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_api_oms_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Order) GetCustomerID() string {
	if x != nil {
		return x.CustomerID
	}
	return ""
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Order) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Order) GetPaymentLink() string {
	if x != nil {
		return x.PaymentLink
	}
	return ""
}

func (x *Order) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Order) GetResultsBase64() []string {
	if x != nil {
		return x.ResultsBase64
	}
	return nil
}

type GetOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderID       string                 `protobuf:"bytes,1,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
	CustomerID    string                 `protobuf:"bytes,2,opt,name=CustomerID,proto3" json:"CustomerID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrderRequest) Reset() {
	*x = GetOrderRequest{}
	mi := &file_api_oms_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderRequest) ProtoMessage() {}

func (x *GetOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_oms_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderRequest.ProtoReflect.Descriptor instead.
func (*GetOrderRequest) Descriptor() ([]byte, []int) {
	return file_api_oms_proto_rawDescGZIP(), []int{1}
}

func (x *GetOrderRequest) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

func (x *GetOrderRequest) GetCustomerID() string {
	if x != nil {
		return x.CustomerID
	}
	return ""
}

type Item struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	ID             string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name           string                 `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	StyleReference string                 `protobuf:"bytes,3,opt,name=StyleReference,proto3" json:"StyleReference,omitempty"`
	PriceID        string                 `protobuf:"bytes,4,opt,name=PriceID,proto3" json:"PriceID,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Item) Reset() {
	*x = Item{}
	mi := &file_api_oms_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_api_oms_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_api_oms_proto_rawDescGZIP(), []int{2}
}

func (x *Item) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetStyleReference() string {
	if x != nil {
		return x.StyleReference
	}
	return ""
}

func (x *Item) GetPriceID() string {
	if x != nil {
		return x.PriceID
	}
	return ""
}

type ItemsWithQuantity struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	ID             string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	StyleReference string                 `protobuf:"bytes,2,opt,name=StyleReference,proto3" json:"StyleReference,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ItemsWithQuantity) Reset() {
	*x = ItemsWithQuantity{}
	mi := &file_api_oms_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ItemsWithQuantity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemsWithQuantity) ProtoMessage() {}

func (x *ItemsWithQuantity) ProtoReflect() protoreflect.Message {
	mi := &file_api_oms_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemsWithQuantity.ProtoReflect.Descriptor instead.
func (*ItemsWithQuantity) Descriptor() ([]byte, []int) {
	return file_api_oms_proto_rawDescGZIP(), []int{3}
}

func (x *ItemsWithQuantity) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *ItemsWithQuantity) GetStyleReference() string {
	if x != nil {
		return x.StyleReference
	}
	return ""
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CustomerID    string                 `protobuf:"bytes,1,opt,name=customerID,proto3" json:"customerID,omitempty"`
	Items         []*ItemsWithQuantity   `protobuf:"bytes,2,rep,name=Items,proto3" json:"Items,omitempty"`
	Image         string                 `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	mi := &file_api_oms_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_oms_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_api_oms_proto_rawDescGZIP(), []int{4}
}

func (x *CreateOrderRequest) GetCustomerID() string {
	if x != nil {
		return x.CustomerID
	}
	return ""
}

func (x *CreateOrderRequest) GetItems() []*ItemsWithQuantity {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CreateOrderRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type CheckIfItemIsInStockRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*ItemsWithQuantity   `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckIfItemIsInStockRequest) Reset() {
	*x = CheckIfItemIsInStockRequest{}
	mi := &file_api_oms_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckIfItemIsInStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckIfItemIsInStockRequest) ProtoMessage() {}

func (x *CheckIfItemIsInStockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_oms_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckIfItemIsInStockRequest.ProtoReflect.Descriptor instead.
func (*CheckIfItemIsInStockRequest) Descriptor() ([]byte, []int) {
	return file_api_oms_proto_rawDescGZIP(), []int{5}
}

func (x *CheckIfItemIsInStockRequest) GetItems() []*ItemsWithQuantity {
	if x != nil {
		return x.Items
	}
	return nil
}

type CheckIfItemIsInStockResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	InStock       bool                   `protobuf:"varint,1,opt,name=InStock,proto3" json:"InStock,omitempty"`
	Items         []*Item                `protobuf:"bytes,2,rep,name=Items,proto3" json:"Items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckIfItemIsInStockResponse) Reset() {
	*x = CheckIfItemIsInStockResponse{}
	mi := &file_api_oms_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckIfItemIsInStockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckIfItemIsInStockResponse) ProtoMessage() {}

func (x *CheckIfItemIsInStockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_oms_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckIfItemIsInStockResponse.ProtoReflect.Descriptor instead.
func (*CheckIfItemIsInStockResponse) Descriptor() ([]byte, []int) {
	return file_api_oms_proto_rawDescGZIP(), []int{6}
}

func (x *CheckIfItemIsInStockResponse) GetInStock() bool {
	if x != nil {
		return x.InStock
	}
	return false
}

func (x *CheckIfItemIsInStockResponse) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type GetItemsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ItemIDs       []string               `protobuf:"bytes,1,rep,name=ItemIDs,proto3" json:"ItemIDs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetItemsRequest) Reset() {
	*x = GetItemsRequest{}
	mi := &file_api_oms_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetItemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemsRequest) ProtoMessage() {}

func (x *GetItemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_oms_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemsRequest.ProtoReflect.Descriptor instead.
func (*GetItemsRequest) Descriptor() ([]byte, []int) {
	return file_api_oms_proto_rawDescGZIP(), []int{7}
}

func (x *GetItemsRequest) GetItemIDs() []string {
	if x != nil {
		return x.ItemIDs
	}
	return nil
}

type GetItemsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*Item                `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetItemsResponse) Reset() {
	*x = GetItemsResponse{}
	mi := &file_api_oms_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetItemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemsResponse) ProtoMessage() {}

func (x *GetItemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_oms_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemsResponse.ProtoReflect.Descriptor instead.
func (*GetItemsResponse) Descriptor() ([]byte, []int) {
	return file_api_oms_proto_rawDescGZIP(), []int{8}
}

func (x *GetItemsResponse) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_api_oms_proto protoreflect.FileDescriptor

const file_api_oms_proto_rawDesc = "" +
	"\n" +
	"\rapi/oms.proto\x12\x03api\"\xcf\x01\n" +
	"\x05Order\x12\x0e\n" +
	"\x02ID\x18\x01 \x01(\tR\x02ID\x12\x1e\n" +
	"\n" +
	"CustomerID\x18\x02 \x01(\tR\n" +
	"CustomerID\x12\x16\n" +
	"\x06Status\x18\x03 \x01(\tR\x06Status\x12\x1f\n" +
	"\x05Items\x18\x04 \x03(\v2\t.api.ItemR\x05Items\x12 \n" +
	"\vPaymentLink\x18\x05 \x01(\tR\vPaymentLink\x12\x14\n" +
	"\x05image\x18\x06 \x01(\tR\x05image\x12%\n" +
	"\x0eresults_base64\x18\a \x03(\tR\rresultsBase64\"K\n" +
	"\x0fGetOrderRequest\x12\x18\n" +
	"\aOrderID\x18\x01 \x01(\tR\aOrderID\x12\x1e\n" +
	"\n" +
	"CustomerID\x18\x02 \x01(\tR\n" +
	"CustomerID\"l\n" +
	"\x04Item\x12\x0e\n" +
	"\x02ID\x18\x01 \x01(\tR\x02ID\x12\x12\n" +
	"\x04Name\x18\x02 \x01(\tR\x04Name\x12&\n" +
	"\x0eStyleReference\x18\x03 \x01(\tR\x0eStyleReference\x12\x18\n" +
	"\aPriceID\x18\x04 \x01(\tR\aPriceID\"K\n" +
	"\x11ItemsWithQuantity\x12\x0e\n" +
	"\x02ID\x18\x01 \x01(\tR\x02ID\x12&\n" +
	"\x0eStyleReference\x18\x02 \x01(\tR\x0eStyleReference\"x\n" +
	"\x12CreateOrderRequest\x12\x1e\n" +
	"\n" +
	"customerID\x18\x01 \x01(\tR\n" +
	"customerID\x12,\n" +
	"\x05Items\x18\x02 \x03(\v2\x16.api.ItemsWithQuantityR\x05Items\x12\x14\n" +
	"\x05image\x18\x03 \x01(\tR\x05image\"K\n" +
	"\x1bCheckIfItemIsInStockRequest\x12,\n" +
	"\x05Items\x18\x01 \x03(\v2\x16.api.ItemsWithQuantityR\x05Items\"Y\n" +
	"\x1cCheckIfItemIsInStockResponse\x12\x18\n" +
	"\aInStock\x18\x01 \x01(\bR\aInStock\x12\x1f\n" +
	"\x05Items\x18\x02 \x03(\v2\t.api.ItemR\x05Items\"+\n" +
	"\x0fGetItemsRequest\x12\x18\n" +
	"\aItemIDs\x18\x01 \x03(\tR\aItemIDs\"3\n" +
	"\x10GetItemsResponse\x12\x1f\n" +
	"\x05Items\x18\x01 \x03(\v2\t.api.ItemR\x05Items2\x97\x01\n" +
	"\fOrderService\x122\n" +
	"\vCreateOrder\x12\x17.api.CreateOrderRequest\x1a\n" +
	".api.Order\x12,\n" +
	"\bGetOrder\x12\x14.api.GetOrderRequest\x1a\n" +
	".api.Order\x12%\n" +
	"\vUpdateOrder\x12\n" +
	".api.Order\x1a\n" +
	".api.Order2\xa4\x01\n" +
	"\fStockService\x12[\n" +
	"\x14CheckIfItemIsInStock\x12 .api.CheckIfItemIsInStockRequest\x1a!.api.CheckIfItemIsInStockResponse\x127\n" +
	"\bGetItems\x12\x14.api.GetItemsRequest\x1a\x15.api.GetItemsResponseB\"Z github.com/sebastvin/commons/apib\x06proto3"

var (
	file_api_oms_proto_rawDescOnce sync.Once
	file_api_oms_proto_rawDescData []byte
)

func file_api_oms_proto_rawDescGZIP() []byte {
	file_api_oms_proto_rawDescOnce.Do(func() {
		file_api_oms_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_oms_proto_rawDesc), len(file_api_oms_proto_rawDesc)))
	})
	return file_api_oms_proto_rawDescData
}

var file_api_oms_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_oms_proto_goTypes = []any{
	(*Order)(nil),                        // 0: api.Order
	(*GetOrderRequest)(nil),              // 1: api.GetOrderRequest
	(*Item)(nil),                         // 2: api.Item
	(*ItemsWithQuantity)(nil),            // 3: api.ItemsWithQuantity
	(*CreateOrderRequest)(nil),           // 4: api.CreateOrderRequest
	(*CheckIfItemIsInStockRequest)(nil),  // 5: api.CheckIfItemIsInStockRequest
	(*CheckIfItemIsInStockResponse)(nil), // 6: api.CheckIfItemIsInStockResponse
	(*GetItemsRequest)(nil),              // 7: api.GetItemsRequest
	(*GetItemsResponse)(nil),             // 8: api.GetItemsResponse
}
var file_api_oms_proto_depIdxs = []int32{
	2,  // 0: api.Order.Items:type_name -> api.Item
	3,  // 1: api.CreateOrderRequest.Items:type_name -> api.ItemsWithQuantity
	3,  // 2: api.CheckIfItemIsInStockRequest.Items:type_name -> api.ItemsWithQuantity
	2,  // 3: api.CheckIfItemIsInStockResponse.Items:type_name -> api.Item
	2,  // 4: api.GetItemsResponse.Items:type_name -> api.Item
	4,  // 5: api.OrderService.CreateOrder:input_type -> api.CreateOrderRequest
	1,  // 6: api.OrderService.GetOrder:input_type -> api.GetOrderRequest
	0,  // 7: api.OrderService.UpdateOrder:input_type -> api.Order
	5,  // 8: api.StockService.CheckIfItemIsInStock:input_type -> api.CheckIfItemIsInStockRequest
	7,  // 9: api.StockService.GetItems:input_type -> api.GetItemsRequest
	0,  // 10: api.OrderService.CreateOrder:output_type -> api.Order
	0,  // 11: api.OrderService.GetOrder:output_type -> api.Order
	0,  // 12: api.OrderService.UpdateOrder:output_type -> api.Order
	6,  // 13: api.StockService.CheckIfItemIsInStock:output_type -> api.CheckIfItemIsInStockResponse
	8,  // 14: api.StockService.GetItems:output_type -> api.GetItemsResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_oms_proto_init() }
func file_api_oms_proto_init() {
	if File_api_oms_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_oms_proto_rawDesc), len(file_api_oms_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_api_oms_proto_goTypes,
		DependencyIndexes: file_api_oms_proto_depIdxs,
		MessageInfos:      file_api_oms_proto_msgTypes,
	}.Build()
	File_api_oms_proto = out.File
	file_api_oms_proto_goTypes = nil
	file_api_oms_proto_depIdxs = nil
}
