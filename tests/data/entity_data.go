package data

import "github.com/braejan/go-bia-challenge/internal/domain/consumption/entity"

func GetConsumptionArrayForTest() (consumptions []*entity.Consumption) {
	// "JUL 03 - JUL 09" 1 1095719.9502499998	1502127.2601399994	1676.8760000000002	1412.6097562610273
	// "JUN 26 - JUL 02" 1 13907652.027070006	796355.9196100001	715.297	1409.2760552337418
	// "JUN 19 - JUN 25" 1 1051010.9760799997	2778538.902079999	1416.4679999999998	1401.5160283531786
	// "JUN 12 - JUN 18" 1 7144698.9787	774140.2560200004	1720.3720000000003	1384.6410152490248
	// "JUN 05 - JUN 11" 1 12955602.134910006	109130.37951999997	2792.68	1371.9244364196995
	// "JUN 01 - JUN 04" 1 559762.1560100004	61442.77255	1196.005	778.5838155823903
	// "JUL 03 - JUL 09" 2 2812578.2076000003	1714762.7678599996	0.0	107.5570928137153
	// "JUN 26 - JUL 02" 2 2717222.7797299977	11488144.13563	0.0	108.10149228491407
	// "JUN 19 - JUN 25" 2 2622722.8854400017	1591995.232300001	0.0	108.77236998495142
	// "JUN 12 - JUN 18" 2 2547738.3204799998	1543837.2360600003	0.0	109.89600306561287
	// "JUN 05 - JUN 11" 2 2441477.4606100004	28248008.731459994	0.0	109.83215558021361
	// "JUN 01 - JUN 04" 2 29631408.80334001	818424.4009800003	0.0	62.96501614154575
	consumptions = append(consumptions, &entity.Consumption{
		Date:               "JUL 03 - JUL 09",
		MeterID:            1,
		ActiveEnergy:       1095719.9502499998,
		ReactiveEnergy:     1502127.2601399994,
		CapacitiveReactive: 1676.8760000000002,
		Solar:              1412.6097562610273,
	})
	consumptions = append(consumptions, &entity.Consumption{
		Date:               "JUN 26 - JUL 02",
		MeterID:            1,
		ActiveEnergy:       13907652.027070006,
		ReactiveEnergy:     796355.9196100001,
		CapacitiveReactive: 715.297,
		Solar:              1409.2760552337418,
	})

	consumptions = append(consumptions, &entity.Consumption{
		Date:               "JUN 19 - JUN 25",
		MeterID:            1,
		ActiveEnergy:       1051010.9760799997,
		ReactiveEnergy:     2778538.902079999,
		CapacitiveReactive: 1416.4679999999998,
		Solar:              1401.5160283531786,
	})

	consumptions = append(consumptions, &entity.Consumption{
		Date:               "JUN 12 - JUN 18",
		MeterID:            1,
		ActiveEnergy:       7144698.9787,
		ReactiveEnergy:     774140.2560200004,
		CapacitiveReactive: 1720.3720000000003,
		Solar:              1384.6410152490248,
	})

	consumptions = append(consumptions, &entity.Consumption{
		Date:               "JUN 05 - JUN 11",
		MeterID:            1,
		ActiveEnergy:       12955602.134910006,
		ReactiveEnergy:     109130.37951999997,
		CapacitiveReactive: 2792.68,
		Solar:              1371.9244364196995,
	})

	consumptions = append(consumptions, &entity.Consumption{
		Date:               "JUN 01 - JUN 04",
		MeterID:            1,
		ActiveEnergy:       559762.1560100004,
		ReactiveEnergy:     61442.77255,
		CapacitiveReactive: 1196.005,
		Solar:              778.5838155823903,
	})

	consumptions = append(consumptions, &entity.Consumption{
		Date:               "JUL 03 - JUL 09",
		MeterID:            2,
		ActiveEnergy:       2812578.2076000003,
		ReactiveEnergy:     1714762.7678599996,
		CapacitiveReactive: 0.0,
		Solar:              107.5570928137153,
	})

	consumptions = append(consumptions, &entity.Consumption{
		Date:               "JUN 26 - JUL 02",
		MeterID:            2,
		ActiveEnergy:       2717222.7797299977,
		ReactiveEnergy:     11488144.13563,
		CapacitiveReactive: 0.0,
		Solar:              108.10149228491407,
	})

	consumptions = append(consumptions, &entity.Consumption{
		Date:               "JUN 19 - JUN 25",
		MeterID:            2,
		ActiveEnergy:       2622722.8854400017,
		ReactiveEnergy:     1591995.232300001,
		CapacitiveReactive: 0.0,
		Solar:              108.77236998495142,
	})

	consumptions = append(consumptions, &entity.Consumption{
		Date:               "JUN 12 - JUN 18",
		MeterID:            2,
		ActiveEnergy:       2547738.3204799998,
		ReactiveEnergy:     1543837.2360600003,
		CapacitiveReactive: 0.0,
		Solar:              109.89600306561287,
	})

	consumptions = append(consumptions, &entity.Consumption{
		Date:               "JUN 05 - JUN 11",
		MeterID:            2,
		ActiveEnergy:       2441477.4606100004,
		ReactiveEnergy:     28248008.731459994,
		CapacitiveReactive: 0.0,
		Solar:              109.83215558021361,
	})

	consumptions = append(consumptions, &entity.Consumption{
		Date:               "JUN 01 - JUN 04",
		MeterID:            2,
		ActiveEnergy:       29631408.80334001,
		ReactiveEnergy:     818424.4009800003,
		CapacitiveReactive: 0.0,
		Solar:              62.96501614154575,
	})
	return
}
