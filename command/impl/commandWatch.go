package impl

// //
// //
// //
// func commandWatch(c *Console, cmd *command.Command) uint {
// 	// if c.context.Selection != context.AreaStatistics {
// 	// 	fmt.Println("Command is not available in this context.")
// 	// 	return 1
// 	// }
//
// 	err := termbox.Init()
//
// 	if err != nil {
// 		output.Println(err)
// 		return 1
// 	}
// 	defer termbox.Close()
//
// 	termbox.SetInputMode(termbox.InputEsc)
//
// 	updateGraph(c, cmd)
//
// 	time.Sleep(10 * time.Second)
//
// 	// update := func(g *ui.BarChart) uint {
// 	// 	ret, err := c.connector.GetStatistic(
// 	// 		c.context.FindSelection(
// 	// 			context.AreaDomain,
// 	// 			context.AreaSubDomain,
// 	// 		).Identifier(),
// 	// 		"5m", listing.DateTimeNow().Add(-40*time.Minute),
// 	// 		listing.DateTimeNow().Add(-10*time.Minute),
// 	// 		[]vo.DatasourceVO{
// 	// 			vo.NewDataSourceVO(vo.DSRequestsCached, vo.TypeHistogram),
// 	// 			vo.NewDataSourceVO(vo.DSRequestsUncached, vo.TypeHistogram),
// 	// 		})
// 	//
// 	// 	if err != nil {
// 	// 		output.Println(err)
// 	// 		return 1
// 	// 	}
// 	//
// 	// 	chartData := []int{}
// 	// 	chartLabel := []string{}
// 	//
// 	// 	for _, val := range ret.Result.GetHistogramValues("requests_cached") {
// 	// 		chartData = append(chartData, int(val)/300)
// 	// 		chartLabel = append(chartLabel, "S1")
// 	// 	}
// 	//
// 	// 	g.Data = chartData
// 	// 	g.DataLabels = chartLabel
// 	//
// 	// 	return 0
// 	// }
//
// 	return 0
// }
//
// func updateGraph(c *Console, cmd *command.Command) uint {
// 	// ret, err := c.connector.GetStatistic(
// 	// 	c.context.FindSelection(
// 	// 		context.AreaDomain,
// 	// 		context.AreaSubDomain,
// 	// 	).Identifier(),
// 	// 	"5m", listing.DateTimeNow().Add(-40*time.Minute),
// 	// 	listing.DateTimeNow().Add(-10*time.Minute),
// 	// 	[]vo.DatasourceVO{
// 	// 		vo.NewDataSourceVO(vo.DSRequestsCached, vo.TypeHistogram),
// 	// 		vo.NewDataSourceVO(vo.DSRequestsUncached, vo.TypeHistogram),
// 	// 	})
// 	//
// 	// if err != nil {
// 	// 	output.Println(err)
// 	// 	return 1
// 	// }
//
// 	maxWidth, maxHeight := termbox.Size()
//
// 	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
//
// 	for y := 0; y < maxHeight; y++ {
// 		for x := 0; x < maxWidth; x++ {
// 			termbox.SetCell(x, y, '█', termbox.ColorDefault, termbox.ColorDefault)
// 		}
// 	}
// 	termbox.Flush()
//
// 	return 0
// }
