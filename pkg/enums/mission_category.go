package enums

type MissionCategory string

const (
	MissionCategoryUnknown       MissionCategory = "unknown"
	MissionCategorySD            MissionCategory = "SD"
	MissionCategoryJP            MissionCategory = "JP"
	MissionCategoryWT            MissionCategory = "WT"
	MissionCategoryJPDK          MissionCategory = "JP_DK"
	MissionCategorySSH           MissionCategory = "SSH"
	MissionCategorySDTOMATO      MissionCategory = "SD_TOMATO"
	MissionCategorySDCMD         MissionCategory = "SD_CMD"
	MissionCategorySDTian        MissionCategory = "SD_TIAN"
	MissionCategorySDBingo       MissionCategory = "SD_BINGO"
	MissionCategoryFooocus       MissionCategory = "FOOOCUS"
	MissionCategoryFooocusLanQue MissionCategory = "FOOOCUS_LAN_QUE"
	MissionCategoryFooocusShaApi MissionCategory = "FOOOCUS_SHA_API"
	MissionCategoryTabby         MissionCategory = "TABBY"
	MissionCategoryOllama        MissionCategory = "OLLAMA"
	MissionCategoryJpConda       MissionCategory = "JP_CONDA"
	MissionCategorySDCat         MissionCategory = "SD_CAT"
	MissionCategorySDFire        MissionCategory = "SD_FIRE"
	MissionCategoryComfyui       MissionCategory = "COMFYUI"
	MissionCategorySDXL          MissionCategory = "SD_XL"
	MissionCategorySDChick       MissionCategory = "SD_CHICK"
	MissionCategoryAscend        MissionCategory = "ASCEND"
	MissionCategorySDWuShan      MissionCategory = "SD_WU_SHAN"
	MissionCategorySDLang        MissionCategory = "SD_LANG"
	MissionCategoryComfyuiKe     MissionCategory = "COMFYUI_KE"
	MissionCategoryChatchat      MissionCategory = "CHATCHAT"
	MissionCategoryLora          MissionCategory = "LORA"
	MissionCategoryFooocusWu     MissionCategory = "FOOOCUS_WU"
	MissionCategorySvdBack       MissionCategory = "SVD_BACK"
	MissionCategorySDJi          MissionCategory = "SD_JI"
	MissionCategorySDShangJin    MissionCategory = "SD_SHANG_JIN"
	MissionCategorySDXiaoChun    MissionCategory = "SD_XIAO_CHUN"
	MissionCategoryComfyuiWu     MissionCategory = "COMFYUI_WU"
	MissionCategoryComfyuiLiu    MissionCategory = "COMFYUI_LIU"
	MissionCategoryComfyuiLi     MissionCategory = "COMFYUI_LI"
	MissionCategoryComfyuiNenly  MissionCategory = "COMFYUI_NENLY"
	MissionCategoryWaiting       MissionCategory = "WAITING"
)

func (obj MissionCategory) Values() []string {
	return []string{
		string(MissionCategoryUnknown),
		string(MissionCategorySD),
		string(MissionCategoryJP),
		string(MissionCategoryWT),
		string(MissionCategoryJPDK),
		string(MissionCategorySSH),
		string(MissionCategorySDTOMATO),
		string(MissionCategorySDCMD),
		string(MissionCategorySDTian),
		string(MissionCategorySDBingo),
		string(MissionCategoryFooocus),
		string(MissionCategoryFooocusLanQue),
		string(MissionCategoryFooocusShaApi),
		string(MissionCategoryTabby),
		string(MissionCategoryOllama),
		string(MissionCategoryJpConda),
		string(MissionCategorySDCat),
		string(MissionCategorySDFire),
		string(MissionCategoryComfyui),
		string(MissionCategorySDXL),
		string(MissionCategorySDChick),
		string(MissionCategoryAscend),
		string(MissionCategorySDWuShan),
		string(MissionCategorySDLang),
		string(MissionCategoryComfyuiKe),
		string(MissionCategoryChatchat),
		string(MissionCategoryLora),
		string(MissionCategoryFooocusWu),
		string(MissionCategorySvdBack),
		string(MissionCategorySDJi),
		string(MissionCategorySDShangJin),
		string(MissionCategorySDXiaoChun),
		string(MissionCategoryComfyuiWu),
		string(MissionCategoryComfyuiLiu),
		string(MissionCategoryComfyuiLi),
		string(MissionCategoryComfyuiNenly),
		string(MissionCategoryWaiting),
	}
}

func (obj MissionCategory) Ptr() *MissionCategory {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
