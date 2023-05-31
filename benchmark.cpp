#include <benchmark/benchmark.h>
#include <array>
#include <memory>

class CustomMemoryManager: public benchmark::MemoryManager {
public:

    int64_t num_allocs;
    int64_t max_bytes_used;


    void Start() BENCHMARK_OVERRIDE {
        num_allocs = 0;
        max_bytes_used = 0;
    }

    void Stop(Result& result) BENCHMARK_OVERRIDE {
        result.num_allocs = num_allocs;
        result.max_bytes_used = max_bytes_used;
    }
};

std::unique_ptr<CustomMemoryManager> mm(new CustomMemoryManager());

#ifdef MEMORY_PROFILER
void *custom_malloc(size_t size) {
    void *p = malloc(size);
    mm.get()->num_allocs += 1;
    mm.get()->max_bytes_used += size;
    return p;
}
#define malloc(size) custom_malloc(size)
#endif
 
constexpr int len = 6;
 
// constexpr function具有inline属性，你应该把它放在头文件中
constexpr auto my_pow(const int i)
{
    return i * i;
}
 
// 使用operator[]读取元素，依次存入1-6的平方
static void bench_array_operator(benchmark::State& state)
{
    std::array<int, len> arr;
    constexpr int i = 1;
    for (auto _: state) {
        arr[0] = my_pow(i);
        arr[1] = my_pow(i+1);
        arr[2] = my_pow(i+2);
        arr[3] = my_pow(i+3);
        arr[4] = my_pow(i+4);
        arr[5] = my_pow(i+5);
    }
}
BENCHMARK(bench_array_operator);
 
// 使用at()读取元素，依次存入1-6的平方
static void bench_array_at(benchmark::State& state)
{
    std::array<int, len> arr;
    constexpr int i = 1;
    for (auto _: state) {
        arr.at(0) = my_pow(i);
        arr.at(1) = my_pow(i+1);
        arr.at(2) = my_pow(i+2);
        arr.at(3) = my_pow(i+3);
        arr.at(4) = my_pow(i+4);
        arr.at(5) = my_pow(i+5);
    }
}
BENCHMARK(bench_array_at);
 
// std::get<>(array)是一个constexpr function，它会返回容器内元素的引用，并在编译期检查数组的索引是否正确
static void bench_array_get(benchmark::State& state)
{
    std::array<int, len> arr;
    constexpr int i = 1;
    for (auto _: state) {
        std::get<0>(arr) = my_pow(i);
        std::get<1>(arr) = my_pow(i+1);
        std::get<2>(arr) = my_pow(i+2);
        std::get<3>(arr) = my_pow(i+3);
        std::get<4>(arr) = my_pow(i+4);
        std::get<5>(arr) = my_pow(i+5);
    }
}
BENCHMARK(bench_array_get);

static void BM_memory(benchmark::State& state) {
    for (auto _ : state)
        for (int i =0; i < 10; i++) {
            benchmark::DoNotOptimize((int *) malloc(100000000));
        }
}

BENCHMARK(BM_memory)->Unit(benchmark::kMillisecond)->Iterations(17);
 
int main(int argc, char** argv)
{
    ::benchmark::RegisterMemoryManager(mm.get());
    ::benchmark::Initialize(&argc, argv);
    ::benchmark::RunSpecifiedBenchmarks();
    ::benchmark::RegisterMemoryManager(nullptr);
}