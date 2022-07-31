This program is a certificates achievement generator that produces PDFs and HTML files.

It allow you to specify following parameters with CLI:

- file: `/home/user/me/file.csv` _(.CSV file path from which list of information will be extracted to generate corresponding certificates)_
- outputType: `pdf` or `html` _(type of file generated: .PDF or .HTML)_

To run it, type the following command:

Compile sources:
````
go build
````

Run the program, (example):
````
./gencert -type pdf -file student.csv
````

Example of generated files:
[![Screenshot-2022-07-31-at-12-23-02.png](https://i.postimg.cc/668DpfZy/Screenshot-2022-07-31-at-12-23-02.png)](https://postimg.cc/qgdmmnjT)
