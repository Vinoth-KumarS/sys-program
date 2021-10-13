#! /bin/bash
#

echo "day03 file - Variable Declaration"


a =10
#echo $a

#simple varibale script 
#
#echo "exporting value from the main data and getting output from the secondaryfile"
name="Vinoth Kumar"
collegeName="Apollo Engineering College"
departmentName="Mechanical Engineering"
place="Chennai"

export name
export collegeName
export departmentName
export place

sh /home/provilty/workspace/sys-program/shellscript/day_03_tst2.sh

#echo "What is your Name : $name"
#echo "What is your College Name : $collegeName"
#echo "Which department you belongs to : $departmentName"
#echo "Which place your College Belongs to : $place"


