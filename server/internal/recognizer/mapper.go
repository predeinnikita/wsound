package recognizer

func recognizeResponseToRecognizeResult(response RecognizeResponse) RecognizeResult {
	var result []FileResult

	for filename, intervals := range response {
		fileResult := FileResult{
			Filename:  filename,
			IsWolf:    len(intervals) > 0,
			Intervals: []Interval{},
		}

		for _, interval := range intervals {
			if len(interval) == 2 {
				fileResult.Intervals = append(fileResult.Intervals, Interval{
					Start: interval[0],
					End:   interval[1],
				})
			}
		}

		result = append(result, fileResult)
	}

	return RecognizeResult{Result: result}
}
