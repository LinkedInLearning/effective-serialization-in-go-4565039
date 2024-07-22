# Effective Serialization In Go: JSON, Protocol Buffers and More
This is the repository for the LinkedIn Learning course Effective Serialization In Go: JSON, Protocol Buffers and More. The full course is available from [LinkedIn Learning][lil-course-url].

![lil-thumbnail-url]

<p>Serialization—taking data structures from the language and converting them to a sequence of bytes—is something every developer encounters at one point or another. In this course, instructor Miki Tebeka teaches you how to work effectively with serialization in Go. This course is targeted to advanced Go developers, and starts with a look at some serialization best practices. Miki then looks at some Go-specific formats as well as JSON. He also shows you how to work with protocol buffers—a powerful binary format—and finishes with a tour of some other popular serialization formats like XML, CSV, and SQL. Each chapter ends with a challenge, so you can test your serialization skills as you go.
 </p>
<p>This course is integrated with GitHub Codespaces, an instant cloud developer environment that offers all the functionality of your favorite IDE without the need for any local machine setup. With GitHub Codespaces, you can get hands-on practice from any machine, at any time—all while using a tool that you’ll likely encounter in the workplace. Check out the “Working with Codespaces” video to learn how to get started.</p>

_See the readme file in the main branch for updated instructions and information._
## Instructions
This repository has branches for each of the videos in the course. You can use the branch pop up menu in github to switch to a specific branch and take a look at the course at that stage, or you can add `/tree/BRANCH_NAME` to the URL to go to the branch you want to access.

## Branches
The branches are structured to correspond to the videos in the course. The naming convention is `CHAPTER#_MOVIE#`. As an example, the branch named `02_03` corresponds to the second chapter and the third video in that chapter. 
Some branches will have a beginning and an end state. These are marked with the letters `b` for "beginning" and `e` for "end". The `b` branch contains the code as it is at the beginning of the movie. The `e` branch contains the code as it is at the end of the movie. The `main` branch holds the final state of the code when in the course.

When switching from one exercise files branch to the next after making changes to the files, you may get a message like this:

    error: Your local changes to the following files would be overwritten by checkout:        [files]
    Please commit your changes or stash them before you switch branches.
    Aborting

To resolve this issue:
	
    Add changes to git using this command: git add .
	Commit changes using this command: git commit -m "some message"

## Installing
1. To use these exercise files, you must have the following installed:
	- [Go SDK](https://go.dev/dl/)
    - [protoc](https://protobuf.dev/downloads/) compiler
    - An IDE such as [VSCode](https://code.visualstudio.com/) or [GoLand](https://www.jetbrains.com/go).
2. Clone this repository into your local machine using the terminal (Mac), CMD (Windows), or a GUI tool like SourceTree.

### Instructor

Miki Tebeka

CEO at 353Solutions
                            

Check out my other courses on [LinkedIn Learning](https://www.linkedin.com/learning/instructors/miki-tebeka?u=104).

[0]: # (Replace these placeholder URLs with actual course URLs)

[lil-course-url]: https://www.linkedin.com/learning/effective-serialization-in-go-json-protocol-buffers-and-more
[lil-thumbnail-url]: https://media.licdn.com/dms/image/D560DAQGMn2FjYJZD5Q/learning-public-crop_675_1200/0/1719254839123?e=2147483647&v=beta&t=Ivaxf0Im9GSnlh4gVOtA__jXZzKcaZNqIwPT9Iv6RqY

