#include <stdio.h>

typedef enum { true, T } boolean;

int main()
{
    int len = 8, depth = 6;
    char treasure[depth][len];

    for(int i=0; i<depth; i++){
        for(int j=0; j<len; j++){
            if(i==0 || i==depth-1 || j==0 || j==len-1){
                treasure[i][j] = '#';
                printf("#");
            } else if (i==2 && (j>1 && j<5) || (i==3 && j==4) || (i==3 && j==6) || (i==4 && j==2)) {
                treasure[i][j] = '#';
                printf("#");
            } else if(i==4 && j==1) {
                treasure[i][j] = 'X';
                printf("X");
            } else {
                treasure[i][j] = '.';
                printf(".");
            }
        }
        printf("\n");
    }

    int y=4, x=1;
    boolean keepMoving = true;
    int arrX[48], arrY[48];
    int ctArrx = 0, ctArry = 0;
    printf("Possible treasures :\n");
    
    do {
        //move X to top
        int tempY = y-1;
        if(tempY > 0){
            //loop X to right
            for(int i=x+1; i<len-1; i++){
                if(treasure[tempY][i] == '.'){
                    //move X to bot
                    int nextY = tempY+1;
                    //keep move to bottom while (x,y) is . and > batas bawah
                    while(nextY<(depth-1)){

                        if(treasure[nextY][i] == '.'){
                            printf("%d,%d \n", nextY, i);
                            arrX[ctArrx] = i; ctArrx++;
                            arrY[ctArry] = nextY; ctArry++;
                        } else {
                            break;
                        }
                        nextY++;
                    }
                } else {
                    break;
                }
            }
        }
        y--;
    } while(y>0);
    printf("\nCompleted treasure :\n");
    for(int i=0; i<depth; i++){
        for(int j=0; j<len; j++){
            if((i==arrY[0] && j==arrX[0]) || (i==arrY[1] && j==arrX[1]) || (i==arrY[2] && j==arrX[2]) || (i==arrY[3] && j==arrX[3]) || (i==arrY[4] && j==arrX[4])){
                printf("$");
            } else {
                printf("%c", treasure[i][j]);     
            }
        }
        printf("\n");
    }
    return 0;
}