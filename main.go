package main

/** Initialize your data structure here. */
func Constructor() Logger {
	return Logger{nil, nil, 0, Set{}}
}

/** Returns true if the message should be printed in the given timestamp, otherwise returns false.
  If this method returns false, the message will not be printed.
  The timestamp is in seconds granularity. */
func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	m := &Message{timestamp: timestamp, message: message}

	// dequeue and remove from set all Expired messages
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

	// if the message is new, enqueue it and add it to the set
	if !this.set.contains(message) {
		this.Enqueue(m)
		this.set.add(message)
		return true
	}
	return false
}

// LinkedList

type Message struct {
	next      *Message
	timestamp int
	message   string
}

type Logger struct {
	first *Message
	last  *Message
	count int
	set   Set
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
func (q *Logger) Dequeue() (i *Message) {
	if q.count == 0 {
		return nil
	}

	item := q.first
	q.first = q.first.next
	q.count--
	return item
}

// SET

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
