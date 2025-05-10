// stands for config
//
//	cfgs.GetString('dto.case')
//
//	cfg := cfgs.New('dto')
//	cfg.Get('case')
//	cfg.Get('model.name')
//	cfg.Get('model.fields')
//
//	cfg2 := cfgd.New('dto.model')
//	cfg2.Get('name')
//	cfg2.Get('fields')
//
// file: 'dto.json'
//
// ```
//
//		{
//			"case":	"pascal_case",
//			"model": {
//	        	"name": "Go",
//	        	"fields": ["Go"]
//	    	}
//		}
//
// ```
//
// 9/May/2024
package cfgs
