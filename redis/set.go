package redis

func Set(str string) error{

	if err:=rdb.Set(ctx,"1",str,0).Err(); err!=nil {
			return err
		}
		return nil

}
