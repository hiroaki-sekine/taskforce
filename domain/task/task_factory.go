package task

type Factory interface {
	New(id int, title, description string, state State) Task
	Start(task Task) Task
	Stop(task Task) Task
	Pause(task Task) Task
	Complete(task Task) Task
	Close(task Task) Task
	GetStateFactory() StateFactory
}

type taskFactory struct {
	stateFactory StateFactory
}

func NewTaskFactory(stateFactory StateFactory) Factory {
	return &taskFactory{
		stateFactory,
	}
}

func (tf *taskFactory) New(id int, title, description string, state State) Task {
	var task Task = &task{
		id:          id,
		title:       title,
		description: description,
		state:       state,
	}
	return task
}

func (tf *taskFactory) Start(task Task) Task {
	task.setState(tf.stateFactory.Doing())
	return task
}

func (tf *taskFactory) Stop(task Task) Task {
	task.setState(tf.stateFactory.Todo())
	return task
}

func (tf *taskFactory) Pause(task Task) Task {
	task.setState(tf.stateFactory.Paused())
	return task
}

func (tf *taskFactory) Complete(task Task) Task {
	task.setState(tf.stateFactory.Completed())
	return task
}

func (tf *taskFactory) Close(task Task) Task {
	task.setState(tf.stateFactory.Closed())
	return task
}

func (tf *taskFactory) GetStateFactory() StateFactory {
	return tf.stateFactory
}
