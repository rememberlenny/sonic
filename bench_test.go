package sonic

import (
	"testing"

	"github.com/bytedance/sonic/cmd"
)

func BenchmarkUnmarshalStructAndInterface(b *testing.B) {

	b.Run("struct", func(b *testing.B) {
		var obj = new(TwitterStruct)
		b.SetBytes(int64(len(TwitterJson)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Unmarshal([]byte(TwitterJson), obj)
		}
	})

	b.Run("interface", func(b *testing.B) {
		var obj = map[string]interface{}{}
		b.SetBytes(int64(len(TwitterJson)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Unmarshal([]byte(TwitterJson), obj)
		}
	})

}

func BenchmarkUnmarshalAndGet(b *testing.B) {

	b.Run("unmarshal", func(b *testing.B) {
		var obj = new(TwitterStruct)
		b.SetBytes(int64(len(TwitterJson)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Unmarshal([]byte(TwitterJson), obj)
			_ = obj.Statuses[0].ID
		}
	})

	b.Run("get", func(b *testing.B) {
		b.SetBytes(int64(len(TwitterJson)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			n, _ := Get([]byte(TwitterJson), "statuses", 0, "id")
			n.Int64()
		}
	})

}

func BenchmarkParseAndInterface(b *testing.B) {

	b.Run("Parse", func(b *testing.B) {
		b.SetBytes(int64(len(TwitterJson)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			root, _ := Get([]byte(TwitterJson), "statuses", 0, "user")
			_ = root.Get("name")
			entities := root.Get("entities")
			_ = entities.Get("url").Get("urls").Index(0).Get("expanded_url")
		}
	})

	b.Run("Interface", func(b *testing.B) {
		b.SetBytes(int64(len(TwitterJson)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			root, _ := Get([]byte(TwitterJson), "statuses", 0, "user")
			user := root.Interface().(map[string]interface{})
			_ = user["name"]
			entities := user["entities"].(map[string]interface{})
			_ = entities["url"].(map[string]interface{})["urls"].([]interface{})[0].(map[string]interface{})["expanded_url"]
		}
	})

}

func BenchmarkGetGetAndGetParse(b *testing.B) {

	b.Run("Get+Get", func(b *testing.B) {
		b.SetBytes(int64(len(TwitterJson)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = Get([]byte(TwitterJson), "statuses", 2, "user", "name")
			_, _ = Get([]byte(TwitterJson), "statuses", 2, "user", "entities")
		}
	})

	b.Run("Get+Parse", func(b *testing.B) {
		b.SetBytes(int64(len(TwitterJson)))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			root, _ := Get([]byte(TwitterJson), "statuses", 2, "user")
			_ = root.Get("name")
			_ = root.Get("entities")
		}
	})
}

func BenchmarkCAndNative(b *testing.B) {
	s := cmd.GenS(5000)
	b.Run("C", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cmd.RunC(s)
		}
	})
	b.Run("native", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cmd.RunNative(s)
		}
	})
	b.Run("go", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			cmd.RunGo(s)
		}
	})
}
