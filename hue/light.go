package hue

type Light struct {
	Name string `json:name`
}

func (b *Bridge) Lights() (map[string]Light, error) {
	// resp, err := b.httpGET("/lights")
	// if err != nil {
	// 	return nil, err
	// }

	// klog.Info(resp.Status)
	// buffered := bufio.NewReader(resp.Body)
	// c0, err := buffered.Peek(1)
	// if err != nil {
	// 	return nil, err
	// }

	// var lights map[string]Light

	// var dest interface{}
	// failed := false
	// if c0[0] == '[' {
	// 	failed = true
	// 	dest = []interface{}{}
	// } else {
	// 	dest = &lights
	// }

	// dec := json.NewDecoder(buffered)
	// err = dec.Decode(&dest)
	// if err != nil {
	// 	return nil, err
	// }

	// if failed {
	// 	klog.Error(dest)
	// 	return nil, errs.Internal(nil)
	// }

	//return lights, err
	return nil, nil
}
