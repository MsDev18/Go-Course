package entity

type Question struct {
	ID              uint
	Text            string
	PossibleAnswers []PossibleAnswer
	CorrectAnswerID uint
	Difficulty      QuestionDifficulty
	CategoryID      uint
}

type PossibleAnswer struct {
	ID     uint
	Text   string
	Choise uint8
}

type PossibleAnswerChoise uint8

func (p PossibleAnswerChoise) IsValid() bool {
	if p >= PossibleAnswerA && p <= PossibleAnswerD {
		return true
	}
	return false
}

const (
	PossibleAnswerA PossibleAnswerChoise = iota + 1
	PossibleAnswerB
	PossibleAnswerC
	PossibleAnswerD
)

type QuestionDifficulty uint8

const (
	QuestionDifficultyEasy QuestionDifficulty = iota + 1
	QuestionDefficultyMedium
	QuestionDefficultyHard
)

func (q QuestionDifficulty) IsValid() bool {
	if q >= QuestionDifficultyEasy && q <= QuestionDefficultyHard {
		return true
	}
	return false
}
