package enums

type MissionType string

const (
	MissionTypeUnknown            MissionType = "unknown"
	MissionTypeSdTime             MissionType = "sd_time"
	MissionTypeSdTxt2Img          MissionType = "txt2img"
	MissionTypeSdImg2Img          MissionType = "img2img"
	MissionTypeJpTime             MissionType = "jp_time"
	MissionTypeWtTime             MissionType = "wt_time"
	MissionTypeSdExtraSingleImage MissionType = "extra-single-image"
	MissionTypeSdApi              MissionType = "sd_api"
	MissionTypeKeyPair            MissionType = "key_pair"
	MissionTypeJpDkTime           MissionType = "jp_dk_time"
	MissionTypeSshTime            MissionType = "ssh_time"
	MissionTypeSdTimePlanHour     MissionType = "sd_time_plan_hour"
	MissionTypeSdTimePlanDay      MissionType = "sd_time_plan_day"
	MissionTypeSdTimePlanWeek     MissionType = "sd_time_plan_week"
	MissionTypeSdTimePlanMonth    MissionType = "sd_time_plan_month"
)

func (obj MissionType) Values() []string {
	return []string{
		string(MissionTypeUnknown),
		string(MissionTypeSdTime),
		string(MissionTypeSdTxt2Img),
		string(MissionTypeSdImg2Img),
		string(MissionTypeJpTime),
		string(MissionTypeWtTime),
		string(MissionTypeSdExtraSingleImage),
		string(MissionTypeSdApi),
		string(MissionTypeKeyPair),
		string(MissionTypeJpDkTime),
		string(MissionTypeSshTime),
		string(MissionTypeSdTimePlanHour),
		string(MissionTypeSdTimePlanDay),
		string(MissionTypeSdTimePlanWeek),
		string(MissionTypeSdTimePlanMonth),
	}
}

func (obj MissionType) Ptr() *MissionType {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
