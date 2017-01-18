package util

import "testing"

func TestCounter(t *testing.T) {
	testData := []string{"a", "b", "b", "c", "c", "d", "e", "e", "e", "f"}
	counter := NewCounter()
	for i := range testData {
		counter.Update(testData[i])
	}

	if counter.Total() != len(testData) {
		t.Error("Incorrect total count")
	}

	elems, freqs := counter.Freqs()
	for i := range freqs {
		if elems[i] == "a" && freqs[i] != 1 {
			t.Fail()
		}
		if elems[i] == "b" && freqs[i] != 2 {
			t.Fail()
		}
		if elems[i] == "c" && freqs[i] != 2 {
			t.Fail()
		}
		if elems[i] == "d" && freqs[i] != 1 {
			t.Fail()
		}
		if elems[i] == "e" && freqs[i] != 3 {
			t.Fail()
		}
	}
	elems, probs := counter.Probs()
	for i := range freqs {
		if elems[i] == "a" && probs[i] != 0.1 {
			t.Fail()
		}
		if elems[i] == "b" && probs[i] != 0.2 {
			t.Fail()
		}
		if elems[i] == "c" && probs[i] != 0.2 {
			t.Fail()
		}
		if elems[i] == "d" && probs[i] != 0.1 {
			t.Fail()
		}
		if elems[i] == "e" && probs[i] != 0.3 {
			t.Fail()
		}
	}
	c := counter.Unique()
	if c != 6 {
		t.Error(c)
	}

	var a string
	fn := func(s interface{}) error {
		a += s.(string)
		return nil
	}
	if err := counter.Apply(fn); err != nil {
		t.Fail()
	}
	t.Logf(a)
	if len(a) != 6 {
		t.Fail()
	}
}

func TestPairCounter(t *testing.T) {
	testData1 := []string{"a", "b", "b", "c", "c", "d", "e", "e", "f", "f"}
	testData2 := []string{"o", "p", "p", "q", "q", "r", "s", "t", "t", "v"}
	counter1 := NewCounter()
	counter2 := NewCounter()
	pairCounter := NewPairCounter()
	for i := range testData1 {
		counter1.Update(testData1[i])
		counter2.Update(testData2[i])
		pairCounter.Update(testData1[i], testData2[i])
	}

	if pairCounter.Total() != len(testData1) {
		t.Error("Incorrect total count")
	}
	if pairCounter.Unique() != 8 {
		t.Error("Incorrect unique count")
	}

	e1 := counter1.Entropy()
	e2 := counter2.Entropy()
	je := pairCounter.JointEntropy()
	t.Log(e1)
	t.Log(e2)
	t.Log(je)
	if je < e1 || je < e2 {
		t.Error("Joint entropy must be greater or equal to entropy")
	}
	if e1+e2 < je {
		t.Error("Sum of entropies must be greater or equal to joint entropy")
	}
}
