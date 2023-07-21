# Go Image Processing API

This is an Image Processing API developed using Go. The application provides two filters: grayscale and blur. It also provides two modes of operation: waitgroup and channel. The waitgroup mode is implemented using `sync.WaitGroup`, and the channel mode uses Go channels for task distribution. 

## Structure of the Project

````go
src/project5
├── README.md
├── filter
│   └── filter.go
├── main.go
└── task
    ├── chan.go
    ├── task.go
    └── waitgrp.go

2 directories, 6 files
````

## Setup and Run

Before running the application, ensure you have Go installed (version 1.16+ recommended).

1. Clone the repository:

````bash
git clone https://github.com/username/GoImageProcessingAPI.git
````

2. Navigate into the `src/project5` directory:

````bash
cd src/project5
````

3. Run the server:

````go
go run main.go --src <srcDir> --dst <dstDir> --filter <grayscale/blur> --task <waitgrp/channel> --poolsize <poolSize>
````

This will start processing images from the source directory (`srcDir`) using the filter and task type specified and will write the processed images to the destination directory (`dstDir`). For the channel-based task, you can specify the pool size (`poolSize`).

## Filters

The available filters are:

- Grayscale: Converts an image to grayscale.
- Blur: Applies a blur filter to an image.

## Task Types

The available task types are:

- Waitgrp: Uses `sync.WaitGroup` for handling tasks.
- Channel: Uses Go channels for distributing tasks among workers.
