package main

type Logger struct {
	first *Message
	last  *Message
	count int
	set   Set
}

/** Initialize your data structure here. */
func Constructor() Logger {
	return Logger{nil, nil, 0, Set{}}
}

/** Returns true if the message should be printed in the given timestamp, otherwise returns false.
  If this method returns false, the message will not be printed.
  The timestamp is in seconds granularity. */
func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	m := &Message{timestamp: timestamp, message: message}

	current := this.first
	for i := 0; i < this.count; i++ {
		if current == nil {
			break
		}
		if timestamp-current.timestamp >= 10 {
			v := this.Dequeue()
			this.set.remove(v.message)
			current = current.next
		}
	}

	// if it does not contain the message enqueue it
	if !this.set.contains(message) {
		this.Enqueue(m)
		this.set.add(message)
		return true
	}
	return false
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */

type Message struct {
	next      *Message
	timestamp int
	message   string
}

// go to the end of the line
func (q *Logger) Enqueue(i *Message) {
	if q.count == 0 {
		q.first = i
		q.last = i
		q.count = 1
		return
	}

	end := q.last
	end.next = i
	q.last = i
	q.count++
	return
}

// first person in line
// invalidate expired messages
func (q *Logger) Dequeue() (i *Message) {
	if q.count == 0 {
		return nil
	}

	item := q.first
	q.first = q.first.next
	q.count--
	return item
}

//SET

type Set map[string]struct{}

func (s Set) add(word string) {
	s[word] = struct{}{}
}

func (s Set) remove(word string) {
	delete(s, word)
}

func (s Set) contains(word string) bool {
	_, ok := s[word]
	return ok
}
