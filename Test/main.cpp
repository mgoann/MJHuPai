#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "hu.h"
#include "split.h"

void vSwap(uint8_t & a, uint8_t & b)
{
    uint8_t temp = a;
    a = b;
    b = temp;
}

void vBubbleSort(uint8_t ai[], uint8_t ucSize)
{
    for (uint8_t i = 0; i < ucSize; i++)
    {
        for (uint8_t j = 1; j < ucSize - i; j++)
        {
            if (ai[j - 1] > ai[j])
            {
                vSwap(ai[j - 1], ai[j]);
            }
        }
    }
}

void m_vRandomArray(uint8_t aucArray[], uint8_t ucArraySize)
{
    for (uint8_t i = 0; i < ucArraySize; i++)
    {
        int index = rand() % ucArraySize;
        vSwap(aucArray[i], aucArray[index]);
    }
}

void vGenRandHandCards(Card aucHandCards[MAX_HANDCARD_NUM], Card & aucLaizi)
{
    Card m_aucSysCards[CARD_MAX_COUNT + 1] = { 0 };

    strcat((char *)m_aucSysCards, (const char *)g_CardWangData);
    strcat((char *)m_aucSysCards, (const char *)g_CardTiaoData);
    strcat((char *)m_aucSysCards, (const char *)g_CardTongData);
    strcat((char *)m_aucSysCards, (const char *)g_CardFengData);
    strcat((char *)m_aucSysCards, (const char *)g_CardJianData);

    int m_ucSysCardsNum = strlen((char *)m_aucSysCards);

    m_vRandomArray(m_aucSysCards, m_ucSysCardsNum);

    memcpy(aucHandCards, m_aucSysCards, MAX_HANDCARD_NUM);
    aucLaizi = aucHandCards[0];
}

void vShowHandCards(Card aucHandCards[MAX_HANDCARD_NUM], Card ucLaizi)
{
    vBubbleSort(aucHandCards, MAX_HANDCARD_NUM);

    printf("[[[[ ");
    for (int i = 0; i < MAX_HANDCARD_NUM; i++)
    {
        printf("0x%02x, ", aucHandCards[i]);
    }
    printf("]]]]  ucLaizi = %02x\n", ucLaizi);
}

void vTest1()
{
    for (int i = 0; i < 10000000; i++)
    {
        Card aucHandCards[MAX_HANDCARD_NUM] = { 0 };
        Card ucLaizi = 0;
        vGenRandHandCards(aucHandCards, ucLaizi);
        //printf("i = %d\n", i);
        bool bMyCanHu = bHu(aucHandCards, ucLaizi);
        bool bQiPaiCanHu = split::bHuHasLaizi(aucHandCards, ucLaizi);
        if (bMyCanHu != bQiPaiCanHu)
        {
            static int iDiffCount = 0;
            printf("i = %d, iDiffCount = %d, bMyCanHu = %d, bQiPaiCanHu = %d\n", i, iDiffCount++, bMyCanHu, bQiPaiCanHu);
            vShowHandCards(aucHandCards, ucLaizi);

            if (bQiPaiCanHu)
            {
                getchar();getchar();getchar();getchar();
            }
            
        }
    }
}

int main()
{
    printf("Hello world\n");

    vTest1();

    getchar();getchar();getchar();getchar();

    return 0;
}

