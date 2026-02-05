package application

type StatusUseCase struct {
}

func NewStatusUseCase() *StatusUseCase {
	return &StatusUseCase{}
}

func (uc *StatusUseCase) Execute() string {
	return "operating"
}
