package utils

//func TestLoggerMiddleware(t *testing.T) {
//	type args struct {
//		next http.Handler
//	}
//	tests := []struct {
//		name string
//		args args
//		want http.Handler
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := LoggerMiddleware(tt.args.next); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("LoggerMiddleware() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestLoggerParent(t *testing.T) {
//	tests := []struct {
//		name string
//		want *logrus.Logger
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := LoggerParent(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("LoggerParent() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestLoggerQueryInit(t *testing.T) {
//	type args struct {
//		db *gorm.DB
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			LoggerQueryInit(tt.args.db)
//		})
//	}
//}
