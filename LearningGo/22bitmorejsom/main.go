package main

import(
	"fmt"
	"encoding/json"
)

type Course struct{
	Name string `json:"courseName"`
	Price int `json:"coursePrice"`
	Platform string `json:"coursePlatform"`
	Password string `json:"-"`
	Tags []string `json:"tags,omitempty"`
}

func main(){
	fmt.Println("JSON Creation\n\n");

	//EncodingJson();
	//DecodeJson();
}


func DecodeJson(){
	jsonDataFromWeb := []byte(`
		{
                "courseName": "ReactJS BootCamp",
                "coursePrice": 299,
                "coursePlatform": "LearnCodeOnline.in",
                "tags": ["web-dev","js"]
        }
	`);

	//Struct Form
	var lcoCourses Course;
	json.Unmarshal(jsonDataFromWeb, &lcoCourses);
	fmt.Printf("%#v\n\n\n",lcoCourses);

	//Key Value Pairs				//cuz i dont know what type of valye is..it maybe int/string/array/etc
	var myOnlineCourses map[string]interface{};
	json.Unmarshal(jsonDataFromWeb, &myOnlineCourses);
	fmt.Println(myOnlineCourses,"\n\n");

	//Loop Through it
	// for key, value := range myOnlineCourses{
	// 	fmt.Printf("%v : %v AND type is %T\n",key,value,value);
	// }


	for key, value := range myOnlineCourses {
		switch v := value.(type) {
		// case string:
		// 	fmt.Printf("%v : %v AND type is %T\n", key, v, v)
		// case int:
		// 	fmt.Printf("%v : %v AND type is %T\n", key, v, v)
		case []interface{}:
			fmt.Printf("%v : ", key)
			for _, item := range v {
				fmt.Printf("%v, ", item)
			}
			fmt.Printf("AND type is %T\n",v);
		default:
			fmt.Printf("%v : %v AND type is %T\n", key, v, v)
		}
	}
}


func EncodingJson(){
	lcoCourses := []Course{
		{"ReactJS BootCamp",299,"LearnCodeOnline.in","abc123",[]string{"web-dev","js"}},
		{"NodetJS BootCamp",199,"LearnCodeOnline.in","ddbc123",[]string{"google","js"}},
		{"golang",0,"YouTube.com","omkar#123", nil },
	};

	//Package this data as JSON Data
	
	//finalJson,err := json.Marshal(lcoCourses); // Format changla nahi mhanun commented kela
										
						//struct, prefix string, indent string
						//		, pratyak line adhi te print honar, for json format -> "\t" || ajun stuf sathi use hota
	finalJson,err := json.MarshalIndent(lcoCourses, "", "\t");
	checkErr(err);

	fmt.Printf("%s\nData Type of finalJson : %T",finalJson,finalJson);
}

func checkErr(err error){
	if err != nil{
		panic(err)
	}
}