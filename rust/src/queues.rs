use std::collections::HashMap;
// use chaser::distance;
 
#[derive(Debug, Clone, PartialEq, Eq, Hash)]
struct Job(u32, u32);

#[allow(dead_code)]
fn pd(execution_cache:&mut HashMap<Job, u32>,
      task_list:&mut Vec<Job>,
      run:u32, speed:u32){
    execution_cache.insert(Job(run, speed), 0);
    task_list.insert(0, Job(1,1));
    task_list.insert(0, Job(2,2));
}


#[allow(dead_code)]
fn check_structs(){
    let mut cache:HashMap<Job,u32>=HashMap::new();
    let mut queue:Vec<Job>=Vec::new();

    queue.insert(0, Job(1,1));
    queue.insert(0, Job(2,2));

    dbg!(&queue);
    pd(&mut cache, &mut queue, 1, 1);
    dbg!(cache.contains_key(&Job(1,1)));
    dbg!(cache.contains_key(&Job(1,2)));
    cache.insert(Job(2,3), 0);
}
