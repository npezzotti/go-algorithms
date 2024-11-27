package maximum_profit_in_job_scheduling

import "sort"

type Job struct {
	startTime int
	endTime   int
	profit    int
}

type Jobs []Job

func (jobs Jobs) Len() int           { return len(jobs) }
func (jobs Jobs) Swap(i, j int)      { jobs[i], jobs[j] = jobs[j], jobs[i] }
func (jobs Jobs) Less(i, j int) bool { return jobs[i].startTime < jobs[j].startTime }

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	var jobs Jobs
	for i := 0; i < len(startTime); i++ {
		jobs = append(jobs, Job{
			startTime: startTime[i],
			endTime:   endTime[i],
			profit:    profit[i],
		})
	}

	sort.Sort(jobs)

	cache := make(map[int]int, jobs.Len())
	var dfs func(i int) int
	dfs = func(i int) int {
		if i == jobs.Len() {
			return 0
		}

		if _, ok := cache[i]; ok {
			return cache[i]
		}

		res := dfs(i + 1)
		j := sort.Search(jobs.Len(), func(k int) bool {
			return jobs[i].endTime <= jobs[k].startTime
		})

		res = max(res, jobs[i].profit+dfs(j))
		cache[i] = res

		return res
	}

	return dfs(0)
}
