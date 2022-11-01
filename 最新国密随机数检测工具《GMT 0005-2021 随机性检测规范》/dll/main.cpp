// Snowflake.cpp : 此文件包含 "main" 函数。程序执行将在此处开始并结束。
//

#include <iostream>
#pragma comment(lib,"random64.lib")
#include "random.h"
#include <Windows.h>
#include <vector>
#include <fstream>


DWORD IterFiles(const char* srcPath, std::vector<std::string>& files)
{
	DWORD dwStatus = 0;

	WIN32_FIND_DATA findFileData;

	std::string path = srcPath;
	path += "\\*.*";

	HANDLE hFind = FindFirstFile(path.c_str(), &findFileData);
	if (hFind != INVALID_HANDLE_VALUE)
	{
		do
		{
			if (strcmp(findFileData.cFileName, ".") == 0 || strcmp(findFileData.cFileName, "..") == 0)
			{
				continue;
			}
			else if (findFileData.dwFileAttributes & FILE_ATTRIBUTE_DIRECTORY)
			{
				//dwStatus = IterFiles(srcNewPath, destNewPath);
			}
			else
			{
				std::string fpath = srcPath;
				fpath += "\\";
				fpath.append(findFileData.cFileName);

				files.push_back(fpath);
			}
			if (dwStatus != 0)
			{
				break;
			}
		} while (FindNextFile(hFind, &findFileData));
	}

	return dwStatus;
}

int main(int argc, char ** argv)
{
#if 0
    unsigned char TestMonoBitFrequencyTestSample[] = { 0xcc, 0x15, 0x6c, 0x4c, 0xe0, 0x02, 0x4d, 0x51, 0x13, 0xd6, 0x80, 0xd7, 0xcc, 0xe6, 0xd8, 0xb2 };
    
    GoSlice data;
    data.data = TestMonoBitFrequencyTestSample;
    data.len = sizeof(TestMonoBitFrequencyTestSample);
    data.cap = sizeof(TestMonoBitFrequencyTestSample);
    GoFloat64 dret = MonoBitFrequency(data);

    std::cout << " MonoBitFrequency " << dret <<std::endl;

    GoSlice data1;
    unsigned char TestFrequencyWithinBlockTestSample[] = { 0xc9, 0xf, 0xda, 0xa2, 0x21, 0x68, 0xc2, 0x34, 0xc4, 0xc6, 0x62, 0x8b, 0x80 };
    data1.data = TestFrequencyWithinBlockTestSample;
    data1.len = sizeof(TestFrequencyWithinBlockTestSample);
    data1.cap = sizeof(TestFrequencyWithinBlockTestSample);
    dret = Frequency(data1);
    
    std::cout << " Frequency " << dret << std::endl;

    dret = Poker(data,4);           std::cout << " Poker " << dret << std::endl;
    dret = Overlapping(data, 3);    std::cout << " Overlapping " << dret << std::endl;
    dret = Runs(data);              std::cout << " Runs " << dret << std::endl;
    dret = RunsDistribution(data);  std::cout << " RunsDistribution " << dret << std::endl;
    dret = LongestRun(data);        std::cout << " LongestRun " << dret << std::endl;
    dret = BinaryDerivative(data, 3);   std::cout << " BinaryDerivative " << dret << std::endl;
    dret = Autocorrelation(data, 1);    std::cout << " Autocorrelation " << dret << std::endl;

    //dret = MatrixRank(data);
    
    dret = Cumulative(data1);       std::cout << " Cumulative " << dret << std::endl;
    dret = ApproximateEntropy(data1, 2);    std::cout << " ApproximateEntropy " << dret << std::endl;

    //dret = LinearComplexity(data, 1000);
    //dret = MaurerUniversal(data); 
    dret = DiscreteFourier(data1);      std::cout << " DiscreteFourier " << dret << std::endl;


#else
	if (2 != argc)
		return 0;

	//const char* path = "D:\\Test\\Snowflake\\data";// argv[1];

	const char* path = argv[1];

	std::vector<std::string> files;

	IterFiles(path, files);

	size_t fNum = files.size();

	if (0 == fNum)
	{
		return 0;		
	}


	int retc[25] = { 0 };

	std::cout << " 样本数量 " << fNum << std::endl;
	std::ifstream inF(files[0], std::ios::binary);
	inF.seekg(0, std::ios_base::end);
	int nFileLen = inF.tellg();
	inF.close();

	char* buf = (char*)calloc(nFileLen, 1);
	GoSlice tempdata = { 0 };
	tempdata.data = buf;
	tempdata.len = nFileLen;
	tempdata.cap = nFileLen;

	for (size_t i = 0; i < fNum; i++)
	{		
		std::cout << files[i] << "  " << std::endl;

		std::ifstream inF(files[i], std::ios::binary);
		inF.read(buf, nFileLen);
		inF.close();

		if (0 == MonoBitFrequency(tempdata))
			retc[0]+=1; std::cout <<".";

		if (0 == Frequency(tempdata))
			retc[1] += 1; std::cout << ".";

		if (0 == Poker(tempdata, 4))
			retc[2] += 1; std::cout << ".";

		if (0 == Poker(tempdata, 8))
			retc[3] += 1; std::cout << ".";

		Overlapping_return oret = {1};
		oret = Overlapping(tempdata, 3);
		if (0 == oret.r0)
			retc[4] += 1; std::cout << ".";

		if (0 == oret.r1)
			retc[5] += 1; std::cout << ".";

		oret = Overlapping(tempdata, 5);
		if (0 == oret.r0)
			retc[6] += 1; std::cout << ".";

		if (0 == oret.r1)
			retc[7] += 1; std::cout << ".";		

		if (0 == Runs(tempdata))
			retc[8] += 1; std::cout << ".";

		if (0 == RunsDistribution(tempdata))
			retc[9] += 1; std::cout << ".";

		if (0 == LongestRun(tempdata))
			retc[10] += 1; std::cout << ".";

		if (0 == BinaryDerivative(tempdata, 3))
			retc[11] += 1; std::cout << ".";

		if (0 == BinaryDerivative(tempdata, 7))
			retc[12] += 1; std::cout << ".";

		if (0 == Autocorrelation(tempdata, 1))
			retc[13] += 1; std::cout << ".";

		if (0 == Autocorrelation(tempdata, 2))
			retc[14] += 1; std::cout << ".";

		if (0 == Autocorrelation(tempdata, 8))
			retc[15] += 1; std::cout << ".";

		if (0 == Autocorrelation(tempdata, 16))
			retc[16] += 1; std::cout << ".";

		if (0 == MatrixRank(tempdata))
			retc[17] += 1; std::cout << ".";

		if (0 == Cumulative(tempdata))
			retc[18] += 1; std::cout << ".";

		if (0 == ApproximateEntropy(tempdata, 2))
			retc[19] += 1; std::cout << ".";

		if (0 == ApproximateEntropy(tempdata, 5))
			retc[20] += 1; std::cout << ".";

		if (0 == LinearComplexity(tempdata, 500))
			retc[21] += 1; std::cout << ".";

		if (0 == LinearComplexity(tempdata, 1000))
			retc[22] += 1; std::cout << ".";

		if (0 == MaurerUniversal(tempdata))
			retc[23] += 1; std::cout << ".";

		if (0 == DiscreteFourier(tempdata))
			retc[24] += 1; std::cout << ".";

		std::cout << std::endl;
	}

	if (buf)
		free(buf);

	std::cout << " 检测结束 " << std::endl;
	
	std::string tests[sizeof(retc) / sizeof(int)] = {
		"单比特频数检测," ,
		"块内频数检测 m=10000," ,
		"扑克检测 m=4," ,
		"扑克检测 m=8," ,
		"重叠子序列检测 m=3 P1,",
		"重叠子序列检测 m=3 P2," ,
		"重叠子序列检测 m=5 P1," ,
		"重叠子序列检测 m = 5 P2," ,
		"游程总数检测," ,
		"游程分布检测," ,
		"块内最大游程检测 m=10000," ,
		"二元推导检测 k=3," ,
		"二元推导检测 k=7," ,
		"自相关检测 d=1," ,
		"自相关检测 d=2," ,
		"自相关检测 d=8," ,
		"自相关检测 d=16," ,
		"矩阵秩检测," ,
		"累加和检测," ,
		"近似熵检测 m=2," ,
		"近似熵检测 m=5," ,
		"线性复杂度检测 m=500," ,
		"线性复杂度检测 m=1000," ,
		"Maurer通用统计检测 L=7 Q=1280," ,
		"离散傅里叶检测"
	};
	std::cout << "源数据, ";
	for (size_t i = 0; i < 25; i++)
			std::cout << " " << tests[i];
	std::cout << std::endl;

	std::cout << "总计, ";
	for (size_t i = 0; i < 25; i++)
		std::cout << ", " << retc[i];
	std::cout << std::endl;
	
#endif
    return 0;
}
