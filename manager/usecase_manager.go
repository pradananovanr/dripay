package manager

import "github.com/rizkyfazri23/dripay/usecase"

type UsecaseManager interface {
	GatewayUsecase() usecase.GatewayUsecase
	MemberUsecase() usecase.MemberUsecase
	TransferUsecase() usecase.TransferUsecase
	DepositUsecase() usecase.DepositUsecase
	HistoryUsecase() usecase.HistoryUsecase
}

type usecaseManager struct {
	repoManager RepoManager
}

func (u *usecaseManager) GatewayUsecase() usecase.GatewayUsecase {
	return usecase.NewGatewayUsecase(u.repoManager.GatewayRepo())
}

func (u *usecaseManager) MemberUsecase() usecase.MemberUsecase {
	return usecase.NewMemberUsecase(u.repoManager.MemberRepo())
}

func (u *usecaseManager) TransferUsecase() usecase.TransferUsecase {
	return usecase.NewTransferUsecase(u.repoManager.TransferRepo())
}

func (u *usecaseManager) DepositUsecase() usecase.DepositUsecase {
	return usecase.NewDepositUsecase(u.repoManager.DepositRepo())
}

func (u *usecaseManager) HistoryUsecase() usecase.HistoryUsecase {
	return usecase.NewHistoryUsecase(u.repoManager.HistoryRepo())
}
func NewUsecaseManager(rm RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: rm,
	}
}
