// Code generated by ent, DO NOT EDIT.

package predicate

import (
	"entgo.io/ent/dialect/sql"
)

// Artwork is the predicate function for artwork builders.
type Artwork func(*sql.Selector)

// ArtworkLike is the predicate function for artworklike builders.
type ArtworkLike func(*sql.Selector)

// Bill is the predicate function for bill builders.
type Bill func(*sql.Selector)

// CDKInfo is the predicate function for cdkinfo builders.
type CDKInfo func(*sql.Selector)

// Campaign is the predicate function for campaign builders.
type Campaign func(*sql.Selector)

// CampaignOrder is the predicate function for campaignorder builders.
type CampaignOrder func(*sql.Selector)

// CloudFile is the predicate function for cloudfile builders.
type CloudFile func(*sql.Selector)

// Collect is the predicate function for collect builders.
type Collect func(*sql.Selector)

// CostAccount is the predicate function for costaccount builders.
type CostAccount func(*sql.Selector)

// CostBill is the predicate function for costbill builders.
type CostBill func(*sql.Selector)

// Device is the predicate function for device builders.
type Device func(*sql.Selector)

// DeviceOrErr calls the predicate only if the error is not nit.
func DeviceOrErr(p Device, err error) Device {
	return func(s *sql.Selector) {
		if err != nil {
			s.AddError(err)
			return
		}
		p(s)
	}
}

// DeviceGpuMission is the predicate function for devicegpumission builders.
type DeviceGpuMission func(*sql.Selector)

// DeviceGpuMissionOrErr calls the predicate only if the error is not nit.
func DeviceGpuMissionOrErr(p DeviceGpuMission, err error) DeviceGpuMission {
	return func(s *sql.Selector) {
		if err != nil {
			s.AddError(err)
			return
		}
		p(s)
	}
}

// DeviceOfflineRecord is the predicate function for deviceofflinerecord builders.
type DeviceOfflineRecord func(*sql.Selector)

// DeviceRebootTime is the predicate function for devicereboottime builders.
type DeviceRebootTime func(*sql.Selector)

// DeviceState is the predicate function for devicestate builders.
type DeviceState func(*sql.Selector)

// EarnBill is the predicate function for earnbill builders.
type EarnBill func(*sql.Selector)

// EnumCondition is the predicate function for enumcondition builders.
type EnumCondition func(*sql.Selector)

// EnumMissionStatus is the predicate function for enummissionstatus builders.
type EnumMissionStatus func(*sql.Selector)

// ExtraService is the predicate function for extraservice builders.
type ExtraService func(*sql.Selector)

// ExtraServiceOrder is the predicate function for extraserviceorder builders.
type ExtraServiceOrder func(*sql.Selector)

// ExtraServicePrice is the predicate function for extraserviceprice builders.
type ExtraServicePrice func(*sql.Selector)

// FrpcInfo is the predicate function for frpcinfo builders.
type FrpcInfo func(*sql.Selector)

// FrpsInfo is the predicate function for frpsinfo builders.
type FrpsInfo func(*sql.Selector)

// Gpu is the predicate function for gpu builders.
type Gpu func(*sql.Selector)

// HmacKeyPair is the predicate function for hmackeypair builders.
type HmacKeyPair func(*sql.Selector)

// IncomeManage is the predicate function for incomemanage builders.
type IncomeManage func(*sql.Selector)

// InputLog is the predicate function for inputlog builders.
type InputLog func(*sql.Selector)

// Invite is the predicate function for invite builders.
type Invite func(*sql.Selector)

// LoginRecord is the predicate function for loginrecord builders.
type LoginRecord func(*sql.Selector)

// Lotto is the predicate function for lotto builders.
type Lotto func(*sql.Selector)

// LottoChanceRule is the predicate function for lottochancerule builders.
type LottoChanceRule func(*sql.Selector)

// LottoGetCountRecord is the predicate function for lottogetcountrecord builders.
type LottoGetCountRecord func(*sql.Selector)

// LottoPrize is the predicate function for lottoprize builders.
type LottoPrize func(*sql.Selector)

// LottoRecord is the predicate function for lottorecord builders.
type LottoRecord func(*sql.Selector)

// LottoUserCount is the predicate function for lottousercount builders.
type LottoUserCount func(*sql.Selector)

// Mission is the predicate function for mission builders.
type Mission func(*sql.Selector)

// MissionOrErr calls the predicate only if the error is not nit.
func MissionOrErr(p Mission, err error) Mission {
	return func(s *sql.Selector) {
		if err != nil {
			s.AddError(err)
			return
		}
		p(s)
	}
}

// MissionBatch is the predicate function for missionbatch builders.
type MissionBatch func(*sql.Selector)

// MissionCategory is the predicate function for missioncategory builders.
type MissionCategory func(*sql.Selector)

// MissionConsumeOrder is the predicate function for missionconsumeorder builders.
type MissionConsumeOrder func(*sql.Selector)

// MissionExtraService is the predicate function for missionextraservice builders.
type MissionExtraService func(*sql.Selector)

// MissionFailedFeedback is the predicate function for missionfailedfeedback builders.
type MissionFailedFeedback func(*sql.Selector)

// MissionKeyPair is the predicate function for missionkeypair builders.
type MissionKeyPair func(*sql.Selector)

// MissionKind is the predicate function for missionkind builders.
type MissionKind func(*sql.Selector)

// MissionOrder is the predicate function for missionorder builders.
type MissionOrder func(*sql.Selector)

// MissionProduceOrder is the predicate function for missionproduceorder builders.
type MissionProduceOrder func(*sql.Selector)

// MissionProduction is the predicate function for missionproduction builders.
type MissionProduction func(*sql.Selector)

// OutputLog is the predicate function for outputlog builders.
type OutputLog func(*sql.Selector)

// PlatformAccount is the predicate function for platformaccount builders.
type PlatformAccount func(*sql.Selector)

// Price is the predicate function for price builders.
type Price func(*sql.Selector)

// ProfitAccount is the predicate function for profitaccount builders.
type ProfitAccount func(*sql.Selector)

// ProfitSetting is the predicate function for profitsetting builders.
type ProfitSetting func(*sql.Selector)

// RechargeCampaignRule is the predicate function for rechargecampaignrule builders.
type RechargeCampaignRule func(*sql.Selector)

// RechargeOrder is the predicate function for rechargeorder builders.
type RechargeOrder func(*sql.Selector)

// RenewalAgreement is the predicate function for renewalagreement builders.
type RenewalAgreement func(*sql.Selector)

// Survey is the predicate function for survey builders.
type Survey func(*sql.Selector)

// SurveyAnswer is the predicate function for surveyanswer builders.
type SurveyAnswer func(*sql.Selector)

// SurveyQuestion is the predicate function for surveyquestion builders.
type SurveyQuestion func(*sql.Selector)

// SurveyQuestionOrErr calls the predicate only if the error is not nit.
func SurveyQuestionOrErr(p SurveyQuestion, err error) SurveyQuestion {
	return func(s *sql.Selector) {
		if err != nil {
			s.AddError(err)
			return
		}
		p(s)
	}
}

// SurveyResponse is the predicate function for surveyresponse builders.
type SurveyResponse func(*sql.Selector)

// Symbol is the predicate function for symbol builders.
type Symbol func(*sql.Selector)

// TransferOrder is the predicate function for transferorder builders.
type TransferOrder func(*sql.Selector)

// TroubleDeduct is the predicate function for troublededuct builders.
type TroubleDeduct func(*sql.Selector)

// User is the predicate function for user builders.
type User func(*sql.Selector)

// UserDevice is the predicate function for userdevice builders.
type UserDevice func(*sql.Selector)

// VXAccount is the predicate function for vxaccount builders.
type VXAccount func(*sql.Selector)

// VXSocial is the predicate function for vxsocial builders.
type VXSocial func(*sql.Selector)

// Wallet is the predicate function for wallet builders.
type Wallet func(*sql.Selector)

// WithdrawAccount is the predicate function for withdrawaccount builders.
type WithdrawAccount func(*sql.Selector)

// WithdrawRecord is the predicate function for withdrawrecord builders.
type WithdrawRecord func(*sql.Selector)
