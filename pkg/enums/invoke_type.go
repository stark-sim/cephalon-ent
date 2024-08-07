package enums

type InvokeType string

const (
	UnKnownInvokeType InvokeType = "unknown"
	ApiInvokeType     InvokeType = "api"
	WebInvokeType     InvokeType = "web"
)

func (obj InvokeType) Values() []string {
	return []string{
		string(UnKnownInvokeType),
		string(ApiInvokeType),
		string(WebInvokeType),
	}
}

func (obj InvokeType) Ptr() *InvokeType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
