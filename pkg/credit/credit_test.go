package credit_test

import (
	"fmt"

	"github.com/sorokinkir/bgo2.1/pkg/credit"
)

func ExampleCalculate() { // имя функции - Example + имя проверяемой функции
	fmt.Println(credit.Calculate(1_000_000_00, 36, 20))
	fmt.Println(credit.Calculate(10_000_00, 36, 20))
	// Output:
	// 371635800 3378888800 13378888800
	// 3716300 33786800 133786800
}
