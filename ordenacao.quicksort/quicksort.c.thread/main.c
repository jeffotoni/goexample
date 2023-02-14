// gcc -O3 -o QuickSort main.c -lpthread --static
// gcc -o quicksort main.c -O3 -march=native -mtune=native -lpthread --static
// @jeffotoni

#ifndef _REENTRANT 
#define _REENTRANT 
#endif

/* #define DEBUG */
#include <pthread.h>
#include <stdlib.h>
#include <stdio.h>
#include <stdbool.h>
#include <time.h>
#include <sys/time.h>

#define MAX_SIZE 200000000
#define MAX_WORKERS 100
#define MAX_PIVOTS 10

/* global variables */
struct Pivot
{
    int index;
    int value;
};

const int g_sortCutoff = 100;   /* cut off for forcing normal quicksort execution */

int g_maxSize;                  /* size of array */
int g_arrayData[MAX_SIZE];      
int g_maxWorkers;               /* number of workers */
pthread_t workers[MAX_WORKERS];

struct WorkerData
{
    int id;
    int * start;
    int n;
    int size;
};

struct WorkerData g_workerData[MAX_WORKERS];    /* contains parameter data of each worker */

pthread_attr_t g_attr;

int g_activeWorkers = 0;
pthread_mutex_t g_lock;

double g_startTime;
double g_finalTime;

/* function prototypes */
void initWorkerData();
void generate(int * start, int * end);
bool isSorted();
int compare(const void * a, const void * b);
void * startThread(void * data);
void parallelQuicksort(int * start, int n, int size);
int getPivot(int * start, int n);
int comparePivot(const void * a, const void * b);
void swap(int * a, int * b);
void printArray(int * start, int * end);
double readTimer();

int main(int argc, const char * argv [])
{
    // Read command line arguments
    g_maxSize = (argc > 1)? atoi(argv[1]) : MAX_SIZE;
    g_maxWorkers = (argc > 2)? atoi(argv[2]) : MAX_WORKERS;
    if(g_maxSize > MAX_SIZE) g_maxSize = MAX_SIZE;
    if(g_maxWorkers > MAX_WORKERS) g_maxWorkers = MAX_WORKERS;
    
    /* Initialize worker data */
    initWorkerData();
    

    g_startTime = readTimer();
    /* Generate test data */
    generate(&g_arrayData[0], &g_arrayData[g_maxSize]);
    //printArray(&g_arrayData[0], &g_arrayData[g_maxSize]);
    
    g_finalTime = readTimer() - g_startTime;
    printf("Load vector time is %g sec\n", g_finalTime);

    /* set global thread attributes */
    pthread_attr_init(&g_attr);
    pthread_attr_setscope(&g_attr, PTHREAD_SCOPE_SYSTEM);
    pthread_attr_setdetachstate(&g_attr, PTHREAD_CREATE_JOINABLE);
    
    /* initialize mutex */
    pthread_mutex_init(&g_lock, NULL);
    
    /* Initialize timer and set start time */
    g_startTime = readTimer();
    
    /* Start the recursive parallel quick sort */
    parallelQuicksort(&g_arrayData[0], g_maxSize, sizeof(int));
    
    /* Set end time */
    g_finalTime = readTimer() - g_startTime;
    
    /* Check if the array is sorted */
    if(isSorted(&g_arrayData[0], &g_arrayData[g_maxSize]))
    {
        printf("Array is sorted\n");
    }
    else
    {
        printf("Array is not sorted\n");
    }
    
    printf("The execution time is %g sec\n", g_finalTime);
    //printArray(&g_arrayData[0], &g_arrayData[g_maxSize]);
    
    /* Print array */
    #ifdef DEBUG
    /* printArray(&g_arrayData[0], &g_arrayData[g_maxSize]); */
    #endif
    
    return 0;
}

void initWorkerData()
{
    int i;
    for(i = 0; i < g_maxWorkers; i++)
    {
        g_workerData[i].id = 0;
        g_workerData[i].start = 0;
        g_workerData[i].n = 0;
        g_workerData[i].size = 0;
    }
}

void generate(int * start, int * end)
{
    if(start == end)
        return;
    while(start != end)
    {
        *start = rand() % 100;
        start++;
    }
}

bool isSorted(int * start, int * end)
{
    start++;
    if(start == end)
        return true;
    while(start != end)
    {
        if(*(start - 1) > *(start))
            return false;
        start++;  
    }
    return true;
}

int compare(const void * a, const void * b)
{
  return ( *(int*)a - *(int*)b );
}

/* worker entry point */
void * startThread(void * data)
{
    struct WorkerData * p = (struct WorkerData *) data;
    int id = p->id;

    #ifdef DEBUG
    printf("worker %d (pthread id %lu) has started\n", id, pthread_self());
    #endif
    parallelQuicksort(p->start, p->n, p->size);
    pthread_exit(0);
}

void parallelQuicksort(int * start, int n, int size)
{
    
    if(n < g_sortCutoff)
    {
        qsort(start, n, sizeof(int), compare);
        return;
    }
    
    /* lock the mutex so we can check the worker count */
    pthread_mutex_lock(&g_lock);
    if(g_activeWorkers < g_maxWorkers)
    {
        /* get worker id from counter */
        pthread_t worker = g_activeWorkers;
        g_activeWorkers++;
        pthread_mutex_unlock(&g_lock);
        
        /* get pivot */
        int pivotIndex = getPivot(start, n);
        int right = n - 1;
        /* Move pivot to end */
        swap(&start[pivotIndex], &start[right]);
        int storeIndex = 0;
        int i;
        for(i = 0; i < right; i++)
        {
            if(start[i] < start[right])
            {
                swap(&start[i], &start[storeIndex]);
                storeIndex++;
            }
        }
        swap(&start[storeIndex], &start[right]);
        pivotIndex = storeIndex;
        
        /* split the current array into two sub arrays */
        g_workerData[worker].id = worker;
        g_workerData[worker].start = start + pivotIndex;
        g_workerData[worker].n = n - pivotIndex;
        g_workerData[worker].size = size;
        
        /* Create a sepparate thread to work on one of the sub arrays */
        pthread_create(&workers[worker], &g_attr, startThread, (void *) &g_workerData[worker]);
        
        /* Let this thread work on another sub array */
        parallelQuicksort(start, pivotIndex, size);
        
        /* join and wait for other thread to finish */
        void * status;
        int rc = pthread_join(workers[worker], &status);
        if(rc)
        {
            printf("ERROR; return code from pthread_join() is %d\n", rc);
            exit(-1);
        }
        #ifdef DEBUG
        printf("Main: completed join with worker %ld (pthread id %lu) having a status of %ld\n", 
            worker, workers[worker], (long)status);
        #endif
    }
    else
    {
        pthread_mutex_unlock(&g_lock);
        /* no available threads, do normal quicksort */
        qsort(start, n, sizeof(int), compare);
    }
}

/* select MAX_PIVOTS number of random pivots and choose one of them */
int getPivot(int * start, int n)
{
    /* if n is low skip selecting pivots */
    if(n < 2)
        return 0;
    struct Pivot pivots[MAX_PIVOTS];
    int maxPivots = (MAX_PIVOTS > n)? n : MAX_PIVOTS;
    int i;
    for(i = 0; i < maxPivots; i++)
    {
        int index = rand() % n;
        pivots[i].index = index;
        pivots[i].value = start[index];
    }

    /* sort the pivots */
    qsort(&pivots[0], maxPivots, sizeof(struct Pivot), comparePivot);

    int pivot = pivots[maxPivots / 2].index;
    #ifdef DEBUG
    printf("pivot is %d\n", pivot);
    printf("pivot value is %d\n", start[pivot]);
    #endif
    return pivot;
    
}

/* compare function that sorts values in increasing order */
int comparePivot(const void * a, const void * b)
{
  return (((struct Pivot*)a)->value - ((struct Pivot*)b)->value);
}

/* swaps to integers in place by reference */
void swap(int * a, int * b)
{
    int tmp = *a;
    *a = *b;
    *b = tmp;
}

/* prints array from start to end */
void printArray(int * start, int * end)
{
    if(start == end)
        return;
    while(start != end)
    {
        printf("%d ", *start);
        start++;
    }
    printf("\n");
}

double readTimer()
{
    static bool initialized = false;
    static struct timeval start;
    struct timeval end;
    if( !initialized )
    {
        gettimeofday( &start, NULL );
        initialized = true;
    }
    gettimeofday( &end, NULL );
    return (end.tv_sec - start.tv_sec) + 1.0e-6 * (end.tv_usec - start.tv_usec);
}
