package main

import (
	"fmt"
	"time"
)

// Event 消息
type Event struct {
	Data int64
}

// Observer 观察者，观察特定的消息
type Observer interface {
	//　消息变动通知观察者
	OnNotify(Event)
}

// Notifier 被观察的对象
type Notifier interface {

	// Register　注册观察者
	Register(Observer)
	// Deregister　删除观察者
	Deregister(Observer)
	// Notify　消息通知
	Notify(Event)
}

// eventObserver Observer 实现
type eventObserver struct {
	id int
}

// eventNotifier Notifier　实现
type eventNotifier struct {
	// 用map来存放Observer
	observers map[Observer]struct{}
}

// OnNotify 收到消息变动
func (o *eventObserver) OnNotify(e Event) {
	fmt.Printf("*** 观察者 %d 接受到变动的消息: %d\n", o.id, e.Data)
}

// Register ...
func (o *eventNotifier) Register(l Observer) {
	o.observers[l] = struct{}{}
}

// Deregister ...
func (o *eventNotifier) Deregister(l Observer) {
	delete(o.observers, l)
}

// Notify 把变化的消息发送给所有的观察者
func (o *eventNotifier) Notify(e Event) {
	for i := range o.observers {
		i.OnNotify(e)
	}
}

func main() {
	// 初始化
	n := eventNotifier{
		observers: map[Observer]struct{}{},
	}

	// 注册两个观察者
	n.Register(&eventObserver{id: 1})
	n.Register(&eventObserver{id: 2})

	// 消息通知
	stop := time.NewTimer(10 * time.Second).C
	tick := time.NewTicker(time.Second).C
	for {
		select {
		case <-stop:
			return
		case t := <-tick:
			n.Notify(Event{Data: t.Unix()})
		}
	}
}
