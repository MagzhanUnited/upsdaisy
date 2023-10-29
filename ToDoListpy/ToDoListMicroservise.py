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


# Observer Design Pattern
class ToDoListObserver:
    def update(self, task):
        print(f"Task '{task}' has been updated.")


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

while True:
    # Get user input
    user_input = input("Enter a command (add/complete/delete/list/quit): ")

    if user_input == "quit":
        break

    if user_input == "list":
        tasks = todo_manager.list_tasks()
        print("Current Tasks:")
        for task in tasks:
            print(task)
    elif user_input in ["add", "complete", "delete"]:
        task = input("Enter the task: ")
        command_obj = ToDoCommand(user_input)
        command_obj.execute(task)

# Print the current tasks when quitting
tasks = todo_manager.list_tasks()
print("Current Tasks:")
for task in tasks:
    print(task)
