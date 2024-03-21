interface option {
    value: string,
    label: string
}

class myOption {
    getwebsiteOptions() {
        let websiteOptions: option[]
        websiteOptions = [
            {
                value: 'ips',
                label: 'ips',
            },
            {
                value: 'website',
                label: '网站',
            },
            {
                value: 'title',
                label: '标题',
            },
            {
                value: 'favicon',
                label: 'favicon',
            }
        ];
        return websiteOptions
    }
}


const myoption = new myOption()


export default myoption