import os

# Singleton Design Pattern
class ToDoListManager:
    _instance = None

    def __new__(cls):
        if cls._instance is None:
            cls._instance = super(ToDoListManager, cls).__new__(cls)
            cls._instance.tasks = []
        return cls._instance

    def add_task(self, task):
        self.tasks.append(task)

    def complete_task(self, task):
        if task in self.tasks:
            self.tasks.remove(task)

    def delete_task(self, task):
        if task in self.tasks:
            self.tasks.remove(task)

    def list_tasks(self):
        return self.tasks

    def save_tasks_to_file(self, task):
        with open("list.txt", "w") as file:
            file.write(task + "\n")

    def load_tasks_from_file(self):
        if os.path.exists("list.txt"):
            with open("list.txt", "r") as file:
                self.tasks = [line.strip() for line in file]

# Observer Design Pattern
class ToDoListObserver:
    def update(self, task):
        print(f"Task '{task}' has been updated.")
        todo_manager.save_tasks_to_file(task)

# Memento Design Pattern
class TaskMemento:
    def __init__(self, tasks):
        self.tasks = tasks

class TaskMementoCareTaker:
    def __init__(self):
        self.mementos = []

    def add_memento(self, memento):
        self.mementos.append(memento)

    def get_memento(self, index):
        if 0 <= index < len(self.mementos):
            return self.mementos[index]
        return None

# Command Design Pattern
class ToDoCommand:
    def __init__(self, command):
        self.command = command

    def execute(self, task):
        if self.command == "add":
            todo_manager.add_task(task)
            observer.update(task)
        elif self.command == "complete":
            todo_manager.complete_task(task)
            observer.update(task)
        elif self.command == "delete":
            todo_manager.delete_task(task)
            observer.update(task)

# Create instances
todo_manager = ToDoListManager()
observer = ToDoListObserver()
memento_care_taker = TaskMementoCareTaker()

# Load tasks from file
todo_manager.load_tasks_from_file()

def ToDoService(user_input, task):
    # print("HERE")
    # Get user input
    # user_input = input("Enter a command (add/complete/delete/list/undo/quit): ")

    if user_input == "quit":
        return ""

    if user_input == "list":
        result = ''
        tasks = todo_manager.list_tasks()
        result = "Current Tasks:"
        for ttask in tasks:
            result += ttask
        return result
    elif user_input == "undo":
        index = int(task)
        memento = memento_care_taker.get_memento(index)
        if memento:
            todo_manager.tasks = memento.tasks
            return "State restored."
        else:
            return "Invalid state index."
    elif user_input in ["add", "complete", "delete"]:
        # Create a memento before executing the command
        memento = TaskMemento(list(todo_manager.tasks))
        memento_care_taker.add_memento(memento)

        command_obj = ToDoCommand(user_input)
        command_obj.execute(task)
        return "Task 'somehitng' has been updated."
    return ""

# Save tasks to file when quitting

# print(ToDoService( "list", "something"))