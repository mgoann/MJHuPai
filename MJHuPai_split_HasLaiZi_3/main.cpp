#include <stdio.h>
#include "public.h"
#include "hu.h"

void vTest1()
{
	Card aucHandCards[MAX_HANDCARD_NUM] = {
		0x01, 0x01,
		0x02, 0x02, 0x02,
		0x03, 0x03, 0x03,
		0x04, 0x04, 0x04,
		0x05, 0x05, 0x05,
	};
	bool bCanHu = bHu(aucHandCards, 0x06);
	printf("vTest1 bCanHu = %d\n", bCanHu);
}

void vTest2()
{
	Card aucHandCards[MAX_HANDCARD_NUM] = {
		0x01, 0x01,
		0x02, 0x02, 0x02,
		0x03, 0x03, 0x03,
		0x04, 0x04, 0x04,
		0x05, 0x05, 0x06,
	};
	bool bCanHu = bHu(aucHandCards, 0x06);
	printf("vTest2 bCanHu = %d\n", bCanHu);
}

void vTest3()
{
	Card aucHandCards[MAX_HANDCARD_NUM] = {
		0x01, 0x01, 0x01,
		0x02,
		0x03, 0x03, 0x03, 
		0x04,
		0x06, 0x06, 0x06,
		0x07, 0x08, 0x19,
	};
	bool bCanHu = bHu(aucHandCards, 0x01);
	printf("vTest3 bCanHu = %d\n", bCanHu);
}

void vTest4()
{
    Card aucHandCards[MAX_HANDCARD_NUM] = {
        0x02, 0x04, 0x04, 0x05, 0x06, 0x06, 0x07, 0x08, 0x21, 0x22, 0x23, 0x28, 0x28, 0x29,
    };
    bool bCanHu = bHu(aucHandCards, 0x29);
    printf("vTest4 bCanHu = %d\n", bCanHu);
}


int main()
{
	//vTest1();
	//vTest2();
	//vTest3();
    vTest4();

	return 0;
}