# Nucleome Data

*Nucleome Data* is a GUI Client Application for *Nucleome Browser* users to host bigWig, bigBed and .hic data from local drive or internet, and browsing them in *Nucleome Browser* with other available data.

If you are looking for a command line tool to host data in servers instead of your personal computer, please visit this website [Nucleome Server](https://github.com/nimezhu/cnbData) instead.

## Dependencies

- Requires Chrome/Chromium >= 70 to be installed.

## Quick Start
Please download the example file and the correspoding executable binary file for your computer OS. This example file for demostration contains a web link of hg19 MTA ChIP Seq peaks bigBed file from [Encode Project](https://www.encodeproject.org/). In this case, with only a few clicks, user can browsing these peaks in *Nucleome Browser*.


### Download Input Example File
- [Example Input File](https://vis.nucleome.org/static/ndata/cnb.xlsx)

### Download Binary Executable Program

- [Linux](https://vis.nucleome.org/static/ndata/current/linux/ndata)

- [Windows](https://vis.nucleome.org/static/ndata/current/win64/ndata.exe)

- [MacOS](https://vis.nucleome.org/static/ndata/current/mac/ndata)

### Start Nucleome Data Service

In Mac or Linux, start a terminal and change work directory to where you put the `ndata` file. Start this program with command line below.

`chmod 755 ndata`

`./ndata`

in Windows 

Just double click `ndata.exe`.

Then follow the steps in GUI Application to add input file and start data service.


## Input Excel 
### Format
The input for Nucleome Data is a simplified Excel/Sheets version for [trackHub](https://genome.ucsc.edu/goldenpath/help/hgTrackHubHelp.html) format. 

Two sheets are required for input excel file. 

The first one is “Config”,  which store the configuration variable values. Currently, `root` variable is the only variable needed. It stores the root path for you store all track data files. It is designed for user conveniently migrating data between computers. All the URIs in other sheets will be the relative path to `root` if their URI are not start with `http` or `https`.

![Sheet Config Example](https://nucleome.github.io/image/sheetConfig.png)

The second one is “Index”, which stores the configuration information of all other sheets.

![Sheet Index Example](https://nucleome.github.io/image/sheetIndex.png)

For each track data sheet, if using four columns, the columns name should be “shortLabel” , “uri,metaLink,longLabel”.

If using two columns, the column name could be any string user defined. Just filled in the column index into the fourth and the fifth column accordingly. 

In sheet "Index", those entries which Id starts with “#” will be ignored.Column "Type" is a reserve entry for future data services. Currently, please use "track" in this column. It support bigWig, bigBed and .hic.

#### Simple Name and URI
![Sheet Data Example](https://nucleome.github.io/image/sheetSimpleData.png)

#### With Long Label and Meta Link
![Sheet Data Example](https://nucleome.github.io/image/sheetData4.png)

## Manage data
