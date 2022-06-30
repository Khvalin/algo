
class MaxHeap {
    arr = [];
    count = 0;

    constructor(size) {
        this.arr = Array.from(Array(size));
    }

    LeftChild(x) {
        return 2 * x + 1;
    }

    RightChild(x) {
        return this.LeftChild(x) + 1;
    }

    Parent(x) {
        return (x - 1) >> 1;
    }


    Insert(val) {
        let i = this.count++;
        this.arr[i] = val;
        while (i != 0 && this.arr[this.Parent(i)] < this.arr[i]) {
            const parent = this.Parent(i);

            // Swap child node with parent
            const temp = this.arr[parent];
            this.arr[parent] = this.arr[i];
            this.arr[i] = temp;

            i = parent;
        }
    }

    Heapify(i = 0) {
        if (this.count < 2) {
            return
        }

        while (i < this.count) {
            //For empty or Heap with single element we need not perform any operation

            let largest = i;
            let left = this.LeftChild(i);
            let right = this.RightChild(i);

            if (left < this.count && this.arr[left] > this.arr[i]) {
                largest = left;
            }
            if (right < this.count && this.arr[right] > this.arr[largest]) {
                largest = right;
            }

            if (largest === i) {
                break;
            }

            // Swap root node with larger child
            const temp = this.arr[largest];
            this.arr[largest] = this.arr[i];
            this.arr[i] = temp;

            // Heapify child nodes to re-order subtree and maintain MinHeap integrity
            i = largest;

        }
    }
}



/**
 * @param {number[][]} courses
 * @return {number}
 */
var scheduleCourse = function (courses) {
    courses = courses.sort((a, b) => {
        if (a[1] === b[1]) {
            return a[0] - b[0]
        }

        return a[1] - b[1]
    })

    const heap = new MaxHeap(courses.length);
    let time = 0;
    for (let i = 0; i < courses.length; i++) {
        const [duration, lastDay] = courses[i];

        // if course duration + curr time <= course lastDay to complete course, we can take the curr course
        if (duration + time <= lastDay) {
            // take the course & add it to Priority Queue
            time += duration;
            heap.Insert(duration);
        }
        // duration of curr course is smaller than the largest course we have taken till now
        else if (heap.count > 0 && heap.arr[0] > duration) {
            // remove the course with biggest duration from MaxHeap & add the curr coarse duration
            time -= heap.arr[0] - duration;     // update time to reflect taking shorter duration course
            heap.arr[0] = duration;             // replace the HeapTop with new duration
            heap.Heapify();                  // re-balance the MaxHeap
        }
    }
    return heap.count;
};