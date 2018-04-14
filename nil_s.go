package do

import "fmt"

func fetchFromChoiceTableNil(id int) (data *ChoiceQuestion) {
	if id == 1001 {
		return &ChoiceQuestion{
			BaseQuestion: BaseQuestion{
				QuestionId:      1001,
				QuestionContent: "HELLO",
			},
			Options: []string{"A", "B"},
		}
	}
	return nil
}

func fetchQuestionNil(id int) (interface{}) {
	data1 := fetchFromChoiceTableNil(id) // 根据ID到选择题表中找题目
	return data1
}

func sendData(data interface{}) {
	fmt.Println("发送数据 ..." , data)
}