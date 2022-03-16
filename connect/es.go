package connect

import (
	"fmt"
	"time"

	es_v6 "github.com/olivere/elastic"
)

var ESClient *es_v6.Client

func InitEsClient(ip string, port string) (err error) {
	ESClient, err = NewElasticClient(ip, port)
	return err
}

func NewElasticClient(ip string, port string) (EsClient *es_v6.Client, err error) {
	url := fmt.Sprintf("http://%s:%s", ip, port)
	//设置嗅探时间为1分钟
	EsClient, err = es_v6.NewClient(es_v6.SetURL(url), es_v6.SetSnifferInterval(1*time.Minute))
	if err != nil {
		return EsClient, err
	}
	return EsClient, nil
}

var Mapping = `{"order":0,"index_patterns":["*"],"aliases":{},"settings":{"index":{"max_result_window":"1000000000","number_of_shards":"5","number_of_replicas":"1","analysis":{"analyzer":{"pinyin_analyzer":{"tokenizer":"my_pinyin"}},"tokenizer":{"my_pinyin":{"type":"pinyin","keep_first_letter":false,"keep_full_pinyin":false,"keep_joined_full_pinyin":true,"keep_none_chinese_in_joined_full_pinyin":true,"keep_none_chinese_in_first_letter":true,"none_chinese_pinyin_tokenize":false,"lowercase":true}}}}},"mappings":{"_default_":{"date_detection":true,"dynamic_date_formats":["yyyy-MM-dd'T'HH:mm:ss.SSS Z","Y-M-d H:m:s","Y/M/d H:m:s","Y年M月d日 H:m:s","Y-M-d H:m","Y/M/d H:m","Y年M月d日 H:m","Y-M-d H","Y/M/d H","Y年M月d日 H","Y-M-d","Y/M/d","Y年M月d日"],"dynamic_templates":[{"number":{"mapping":{"store":true},"match_mapping_type":"long","match":"*"}},{"ik_text":{"match":"*_Ik","match_mapping_type":"string","mapping":{"fielddata":true,"analyzer":"ik_max_word","type":"text"}}},{"string_KT":{"match":"*_KeywordIk","match_mapping_type":"string","mapping":{"type":"keyword","store":true,"fielddata":"true","fields":{"text":{"analyzer":"ik_max_word","type":"text","fielddata":true}}}}},{"string_KTP":{"match":"*_KeywordIkPinyin","match_mapping_type":"string","mapping":{"type":"keyword","store":true,"fielddata":false,"fields":{"text":{"analyzer":"ik_max_word","type":"text","fielddata":true},"pinyin":{"analyzer":"pinyin_analyzer","fielddata":true,"type":"text"}}}}},{"string_KP":{"match":"*_KeywordPinyin","match_mapping_type":"string","mapping":{"type":"keyword","store":true,"fielddata":"true","fields":{"pinyin":{"analyzer":"pinyin_analyzer","fielddata":true,"type":"text"}}}}},{"nested":{"match":"*_Nested","mapping":{"type":"nested"}}},{"gps":{"match":"*_Gps","mapping":{"type":"geo_point"}}},{"binary":{"match":"*_Binary","mapping":{"type":"binary"}}},{"no_index":{"match":"*_No","mapping":{"index":false}}},{"keyword":{"match":"*","match_mapping_type":"string","mapping":{"fielddata":false,"store":true,"type":"keyword"}}},{"string_arr_TKP":{"match":"*_ArrKeywordIkPinyin","match_mapping_type":"string","mapping":{"type":"text","analyzer":"ik_max_word","fields":{"keyword":{"type":"keyword","store":true,"ignore_above":256},"pinyin":{"analyzer":"pinyin_analyzer","fielddata":true,"type":"text"}}}}}]}}}`

var IndexMapping = `{
			"properties": {
				"workflow.update_time": {
					"type": "keyword"
				}
			},
            "date_detection": false,
            "dynamic_date_formats": [
               "yyyy-MM-dd'T'HH:mm:ss.SSS Z",
               "Y-M-d H:m:s",
               "Y/M/d H:m:s",
               "Y年M月d日 H:m:s",
               "Y-M-d H:m",
               "Y/M/d H:m",
               "Y年M月d日 H:m",
               "Y-M-d H",
               "Y/M/d H",
               "Y年M月d日 H",
               "Y-M-d",
               "Y/M/d",
               "Y年M月d日"
            ],
            "dynamic_templates": [
                {
                   "float_num": {
                      "mapping": {
                         "type":"float",
                         "store": true
                      },
                      "match_mapping_type": "*",
                      "match": "*__float"
                   }
                },
                {
                   "double_num": {
                      "mapping": {
                          "type":"double",
                         "store": true
                      },
                      "match_mapping_type": "*",
                      "match": "*__double"
                   }
                },
                {
                   "string_time": {
                      "mapping": {
                          "type":"date",
                          "format":"epoch_millis || epoch_second || yyyy-MM-dd HH:mm:ss||yyyy-MM-dd"
                      },
                      "match_mapping_type": "string",
                      "match": "*_time"
                   }
                },
               {
                  "ik_text": {
                     "match": "*_Ik",
                     "match_mapping_type": "string",
                     "mapping": {
                        "fielddata": true,
                        "analyzer": "ik_max_word",
                        "type": "text"
                     }
                  }
               },
               {
                  "string_KT": {
                     "match": "*_KeywordIk",
                     "match_mapping_type": "string",
                     "mapping": {
                        "type": "keyword",
                        "store": true,
                        "fielddata": "true",
                        "fields": {
                           "text": {
                              "analyzer": "ik_max_word",
                              "type": "text",
                              "fielddata": true
                           }
                        }
                     }
                  }
               },
               {
                  "string_KTP": {
                     "match": "*_KeywordIkPinyin",
                     "match_mapping_type": "string",
                     "mapping": {
                        "type": "keyword",
                        "store": true,
                        "fielddata": false,
                        "fields": {
                           "text": {
                              "analyzer": "ik_max_word",
                              "type": "text",
                              "fielddata": true
                           },
                           "pinyin": {
                              "analyzer": "pinyin_analyzer",
                              "fielddata": true,
                              "type": "text"
                           }
                        }
                     }
                  }
               },
               {
                  "string_KP": {
                     "match": "*_KeywordPinyin",
                     "match_mapping_type": "string",
                     "mapping": {
                        "type": "keyword",
                        "store": true,
                        "fielddata": "true",
                        "fields": {
                           "pinyin": {
                              "analyzer": "pinyin_analyzer",
                              "fielddata": true,
                              "type": "text"
                           }
                        }
                     }
                  }
               },
               {
                  "nested": {
                     "match": "*_Nested",
                     "mapping": {
                        "type": "nested"
                     }
                  }
               },
               {
                  "gps": {
                     "match": "*_Gps",
                     "mapping": {
                        "type": "geo_point"
                     }
                  }
               },
               {
                  "binary": {
                     "match": "*_Binary",
                     "mapping": {
                        "type": "binary"
                     }
                  }
               },
               {
                  "no_index": {
                     "match": "*_No",
                     "mapping": {
                        "index": false
                     }
                  }
               },
               {
                  "string_arr_TKP": {
                     "match": "*_ArrKeywordIkPinyin",
                     "match_mapping_type": "string",
                     "mapping": {
                        "type": "text",
                        "analyzer": "ik_max_word",
                        "fields": {
                           "keyword": {
                              "type": "keyword",
                              "store": true,
                              "ignore_above": 256
                           },
                           "pinyin": {
                              "analyzer": "pinyin_analyzer",
                              "fielddata": true,
                              "type": "text"
                           }
                        }
                     }
                  }
               },
            {
                "keyword": {
                    "match": "*",
                    "match_mapping_type": "string",
                    "mapping": {
                    "fielddata": false,
                    "store": true,
                    "type": "keyword"
                    }
                }
            }
            ]
         }`
