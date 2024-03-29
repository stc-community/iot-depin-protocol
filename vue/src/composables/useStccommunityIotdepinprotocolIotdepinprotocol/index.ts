/* eslint-disable @typescript-eslint/no-unused-vars */
import { useQuery, type UseQueryOptions, useInfiniteQuery, type UseInfiniteQueryOptions } from "@tanstack/vue-query";
import { useClient } from '../useClient';
import type { Ref } from 'vue'

export default function useStccommunityIotdepinprotocolIotdepinprotocol() {
  const client = useClient();
  const QueryParams = ( options: any) => {
    const key = { type: 'QueryParams',  };    
    return useQuery([key], () => {
      return  client.StccommunityIotdepinprotocolIotdepinprotocol.query.queryParams().then( res => res.data );
    }, options);
  }
  
  const QueryDevice = (deviceName: string,  options: any) => {
    const key = { type: 'QueryDevice',  deviceName };    
    return useQuery([key], () => {
      const { deviceName } = key
      return  client.StccommunityIotdepinprotocolIotdepinprotocol.query.queryDevice(deviceName).then( res => res.data );
    }, options);
  }
  
  const QueryDeviceAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryDeviceAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.StccommunityIotdepinprotocolIotdepinprotocol.query.queryDeviceAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryKv = (index: string, query: any, options: any) => {
    const key = { type: 'QueryKv',  index, query };    
    return useQuery([key], () => {
      const { index,query } = key
      return  client.StccommunityIotdepinprotocolIotdepinprotocol.query.queryKv(index, query ?? undefined).then( res => res.data );
    }, options);
  }
  
  const QueryKvAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryKvAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.StccommunityIotdepinprotocolIotdepinprotocol.query.queryKvAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryEventPb = (index: string, query: any, options: any) => {
    const key = { type: 'QueryEventPb',  index, query };    
    return useQuery([key], () => {
      const { index,query } = key
      return  client.StccommunityIotdepinprotocolIotdepinprotocol.query.queryEventPb(index, query ?? undefined).then( res => res.data );
    }, options);
  }
  
  const QueryEventPbAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryEventPbAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.StccommunityIotdepinprotocolIotdepinprotocol.query.queryEventPbAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  return {QueryParams,QueryDevice,QueryDeviceAll,QueryKv,QueryKvAll,QueryEventPb,QueryEventPbAll,
  }
}