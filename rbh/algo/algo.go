package algo

// algo.Model
type Model struct {
	// parameters
	para  int
	count int

	// solve
	prices                                 []float64
	linearized                             []float64
	last_dis, curr_dis, last_avg, curr_avg float64
	own                                    bool

	// statistics
	buyprice, algorithm float64 // open price, current price, algorithms
	orders              int
}

// model := algo.NewModel(para)
func NewModel(para int) *Model { return &Model{para: para} }

// shares := model.Solve(price)
func (s *Model) Solve(price float64) int {

	s.prices = append(s.prices, price)

	if len(s.prices) < s.para {
		return 0
	}

	s.linearized = append(s.linearized, Lazy_linear_reg(s.prices, s.para))

	if len(s.prices) < s.para*2-1 {
		return 0
	}

	s.last_dis = s.curr_dis
	s.last_avg = s.curr_avg

	s.curr_dis, s.curr_avg = calculate(s.linearized, s.para)
	// fmt.Printf("%.4f\t%.4f\t%.4f\t%.4f\n", s.last_avg, s.curr_avg, s.last_dis, s.curr_dis)

	if s.curr_avg < 0 && s.last_dis < s.curr_dis && s.own == false {
		s.orders++
		s.buyprice = price + 0.01
		s.own = true
		//		fmt.Printf("%.2f\t%.4f\t%.4f\t%.4f\t%.4f\tbuy\n", s.buyprice, s.last_avg, s.curr_avg, s.last_dis, s.curr_dis)
		return +1
	} else if s.last_avg < s.curr_avg && s.last_dis > s.curr_dis && s.own == true {
		if price-s.buyprice != 0 {
			s.orders++
			s.algorithm += (price - s.buyprice)
			s.own = false
			//			fmt.Printf("%.2f\t%.4f\t%.4f\t%.4f\t%.4f\tsell\n", price, s.last_avg, s.curr_avg, s.last_dis, s.curr_dis)
			return -1
		}
		return 0

	} else {
		//		fmt.Printf("%.2f\t%.4f\t%.4f\t%.4f\t%.4f\n", price, s.last_avg, s.curr_avg, s.last_dis, s.curr_dis)
		return +0
	}
}

func Lazy_linear_reg(price []float64, para int) float64 {

	data := price[len(price)-para:]

	var m float64
	var s int

	for i := 1; i <= para; i++ {
		s += i
	}

	for i := 0; i < para; i++ {
		m = m + ((float64(para)*(float64(i)+1) - float64(s)) * data[i])
	}

	return m
}

func calculate(linearized []float64, para int) (float64, float64) {

	sum := 0.0
	data := linearized[len(linearized)-para:]

	for _, f := range data {
		sum += f
	}

	curr_avg := sum / float64(para)
	curr_dis := curr_avg - data[para-1]

	return curr_dis, curr_avg
}

// algorithm, orders := model.Stats()
func (s *Model) Stats() (float64, int) { return s.algorithm, s.orders }
